package structs

import "fmt"

var teslaModelS = struct {
	Brand, Model, Color string
	weightInKg          int
}{
	Brand:      "Tesla",
	Model:      "Model S",
	Color:      "Red",
	weightInKg: 2250,
}

func PrintAnonymousStruct() {
	fmt.Println(teslaModelS)
}
