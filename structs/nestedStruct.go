package structs

import "fmt"

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Salary  float64
	Address Address // this field is the nested struct Address within Employee
}

func InitAndPrintNestedStruct() {
	homer := Employee{
		Name:    "Homer",
		Age:     39,
		Salary:  724.38,
		Address: Address{City: "Springfield", State: "Oregon"},
	}

	// we can print the nested struct fields using two instances of the '.' operator:
	fmt.Println("City:", homer.Address.City)   // City: Springfield
	fmt.Println("State:", homer.Address.State) // State: Oregon
}
