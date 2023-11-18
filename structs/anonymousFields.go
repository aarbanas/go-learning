package structs

import "fmt"

type Country struct {
	string
	int
}

func InitAndPrintAnonymousFields() {
	var croatia Country

	// this is how we assign values to the anonymous fields,
	// the name of the fields is the same as the data type
	croatia.string = "Croatia"
	croatia.int = 67413000

	fmt.Println("Country name:", croatia.string) // Country name: France
	fmt.Println("Population:", croatia.int)      // Population: 67413000
}
