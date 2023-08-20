package main

import (
	"dimension/pkg/domain"
	"encoding/json"
	"fmt"
)

func main() {

	dim, err := domain.NewDimension(*domain.NewSpherePair(domain.A, domain.Black), *domain.NewSpherePair(domain.B, domain.Black))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dim)

	u, err := json.MarshalIndent(domain.GetGeometry(), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(u))
}
