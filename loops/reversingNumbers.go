package loops

import (
	"fmt"

	. "github.com/aarbanas/go-learning/errors"
)

func ReversNumbers() int {
	var inputNumber int
	result := 0

	fmt.Println("Enter number:")
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		ScanErrorMessage(err)
		return 0
	}

	for inputNumber > 0 {
		lastDecimal := inputNumber % 10

		result = (result * 10) + lastDecimal

		inputNumber /= 10
	}

	return result
}
