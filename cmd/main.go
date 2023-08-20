package main

import (
	"dimension/pkg/domain"
)

func main() {

	//var geomList GeometryList
	//
	////open up the tests file
	////for each test
	////	build an execution, run it, check status code
	//
	//// Read the file
	//data, err := os.ReadFile("data_processing/geometry.json")
	//if err != nil {
	//	log.Fatalf("Error reading file: %v", err)
	//}
	//// Unmarshal JSON data into map[string]interface{}
	//err = json.Unmarshal(data, &geomList)
	//if err != nil {
	//	log.Fatalf("Error unmarshalling JSON data: %v", err)
	//}
	//
	//// Now, geomList.Geometry contains your geometries
	//for _, geom := range geomList.Geometry {
	//	fmt.Printf("ID: %d, PolarAngle: %f, Adjacency: %s\n", geom.ID, geom.PolarAngle, geom.Adjacency)
	//}

	dim, err := domain.NewDimension(domain.NewSphere(domain.Black), domain.NewSphere(domain.Black))

}

//func main() {
//	expression, _ := govaluate.NewEvaluableExpression("(7 && !b)")
//	parameters := map[string]interface{}{"7": true, "b": false}
//	result, _ := expression.Evaluate(parameters)
//	fmt.Println(result) // Outputs: true
//}
