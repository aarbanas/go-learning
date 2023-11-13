package slices

import "fmt"

func CreateSlice() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

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
