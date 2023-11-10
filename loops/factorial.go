package loops

import "fmt"

func CalculateFactorial() int {
	var inputNumber int
	factorial := 1

	fmt.Println("Enter number:")
	fmt.Scan(&inputNumber)
	if inputNumber < 1 {
		fmt.Println("Number must be bigger than 0")
		return 0
	}

	for i := 1; i < inputNumber+1; i++ {
		factorial *= i
	}

	return factorial
}
