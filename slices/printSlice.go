package slices

import "fmt"

func PrintDoubledSliceElements(intSlice []int) {
	fmt.Println("Doubled elements of slice are: ")
	for _, element := range intSlice {
		fmt.Println(element * 2)
	}
}
