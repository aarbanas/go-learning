package factorial

import "fmt"

func Calculate() {
	// Write your code here
	var inputNumber int
	factorial := 1

	fmt.Scan(&inputNumber)
	if inputNumber < 1 {
		fmt.Println("Number must be bigger than 0")
		return
	}

	for i := 1; i < inputNumber+1; i++ {
		factorial *= i
	}

	fmt.Println(factorial)
}
