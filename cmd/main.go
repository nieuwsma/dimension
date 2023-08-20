package main

//func main() {
//
//	//var geomList GeometryList
//	//
//	////open up the tests file
//	////for each test
//	////	build an execution, run it, check status code
//	//
//	//// Read the file
//	//data, err := os.ReadFile("data_processing/geometry.json")
//	//if err != nil {
//	//	log.Fatalf("Error reading file: %v", err)
//	//}
//	//// Unmarshal JSON data into map[string]interface{}
//	//err = json.Unmarshal(data, &geomList)
//	//if err != nil {
//	//	log.Fatalf("Error unmarshalling JSON data: %v", err)
//	//}
//	//
//	//// Now, geomList.Geometry contains your geometries
//	//for _, geom := range geomList.Geometry {
//	//	fmt.Printf("ID: %d, PolarAngle: %f, Adjacency: %s\n", geom.ID, geom.PolarAngle, geom.Adjacency)
//	//}
//
//	dim := NewDimension()
//
//	expression, err := Parse("((7 & 9 & 11) ^ (12 & 10 & 8))")
//	if err != nil {
//		print(err)
//	}
//	dim.Dimension[7] = &Sphere{
//		Green,
//	}
//	dim.Dimension[9] = &Sphere{
//		Green,
//	}
//	dim.Dimension[11] = &Sphere{
//		Green,
//	}
//	result := expression.Evaluate(dim)
//	fmt.Println(result) // Outputs: true
//	dim.Dimension[1] = &Sphere{
//		Green,
//	}
//	result = expression.Evaluate(dim)
//	fmt.Println(result) // Outputs: true
//	dim.Dimension[12] = &Sphere{
//		Green,
//	}
//	dim.Dimension[10] = &Sphere{
//		Green,
//	}
//	dim.Dimension[8] = &Sphere{
//		Green,
//	}
//	result = expression.Evaluate(dim)
//	fmt.Println(result) // Outputs: true
//
//	print(dim.String())
//}

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expression, _ := govaluate.NewEvaluableExpression("(7 && !b)")
	parameters := map[string]interface{}{"7": true, "b": false}
	result, _ := expression.Evaluate(parameters)
	fmt.Println(result) // Outputs: true
}
