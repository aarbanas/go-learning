package slices

import (
	"fmt"

	. "github.com/aarbanas/go-learning/errors"
)

func CreateSlice() {
	var a, b, c int

	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		ScanErrorMessage(err)
		return
	}

	// Create intSlice containing the a, b and c variables below:
	intSlice := []int{a, b, c}

	// Create make function to initialize slice and populate values
	makeSlice := make([]int, 3)
	makeSlice[0] = a
	makeSlice[1] = b
	makeSlice[2] = c

	fmt.Println(intSlice)  // print intSlice here!
	fmt.Println(makeSlice) // print makeSlice here!
}
