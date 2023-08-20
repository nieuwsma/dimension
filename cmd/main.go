package main

import (
	"dimension/pkg/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	dim, err := domain.NewDimension(*domain.NewSpherePair(domain.A, domain.Black), *domain.NewSpherePair(domain.B, domain.Black))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dim)

}

func getGeoList() {
	var geomList domain.Geometries

	//open up the tests file
	//for each test
	//	build an execution, run it, check status code

	// Read the file
	data, err := os.ReadFile("data_processing/geometry.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &geomList)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	// Now, geomList.Geometry contains your geometries
	for _, geom := range geomList.Geometry {
		fmt.Printf("ID: %d, PolarAngle: %f, Neighbors: %s\n", geom.ID, geom.PolarAngle, geom.Neighbors)
	}
}
