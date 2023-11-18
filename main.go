package main

import (
	"fmt"

	. "github.com/aarbanas/go-learning/loops"
	. "github.com/aarbanas/go-learning/slices"
	. "github.com/aarbanas/go-learning/structs"
)

func main() {
	var commandValues int

	fmt.Println("*********************************")
	fmt.Println("Enter one of next commands to do the action: ")
	fmt.Println("1 -> Calculate factorial of provided number")
	fmt.Println("2 -> Reverse numbers order of provided number")
	fmt.Println("3 -> Append slice")
	fmt.Println("4 -> Struct tags")

	fmt.Println("*********************************")

	fmt.Scan(&commandValues)

	switch commandValues {
	case 1:
		res := CalculateFactorial()
		fmt.Println(res)
		break
	case 2:
		res := ReversNumbers()
		fmt.Println(res)
		break
	case 3:
		AppendSlice()
		break
	case 4:
		ExecuteJsonEncoding()
		break
	default:
		fmt.Println("Invalid value provided")
	}

}
