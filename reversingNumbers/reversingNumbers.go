package reversingNumbers

import "fmt"

func Revers() int {
	var inputNumber int
	result := 0

	fmt.Println("Enter number:")
	fmt.Scan(&inputNumber)

	for inputNumber > 0 {
		lastDecimal := inputNumber % 10

		result = (result * 10) + lastDecimal

		inputNumber /= 10
	}

	return result
}
