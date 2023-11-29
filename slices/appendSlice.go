package slices

import (
	"fmt"
	"log"
	"reflect"

	. "github.com/aarbanas/go-learning/errors"
)

func solve(sl1, sl2 []int) []int {
	sl1 = append(sl1, sl2...)

	return sl1
}

func AppendSlice() {
	var num1, num2 int
	fmt.Println("Enter slices length (separate with white space)")
	_, err := fmt.Scanln(&num1, &num2)
	if err != nil {
		ScanErrorMessage(err)
		return
	}

	sl1, sl2 := make([]int, num1), make([]int, num2)

	for i := range sl1 {
		_, err := fmt.Scanln(&sl1[i])
		if err != nil {
			ScanErrorMessage(err)
			return
		}

	}

	for i := range sl2 {
		_, err := fmt.Scanln(&sl2[i])
		if err != nil {
			ScanErrorMessage(err)
			return
		}
	}

	if !reflect.DeepEqual(append(sl1, sl2...), solve(sl1, sl2)) {
		log.Fatal("slices were not properly appended")
	}

	slice := solve(sl1, sl2)
	fmt.Println("Slice: ", slice)
	PrintDoubledSliceElements(slice)
}
