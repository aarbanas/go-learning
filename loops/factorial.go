package loops

import (
	"fmt"

	. "github.com/aarbanas/go-learning/errors"
)

func CalculateFactorial() int {
	var inputNumber int
	factorial := 1

	fmt.Println("Enter number:")
	fmt.Scan(&inputNumber)
	if inputNumber < 1 {
		err := NegativeNumError()
		fmt.Println(err)
		return 0
	}

	for i := 1; i < inputNumber+1; i++ {
		factorial *= i
	}

	return factorial
}
