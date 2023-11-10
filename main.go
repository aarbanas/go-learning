package main

import (
	"fmt"

	"github.com/aarbanas/go-learning/loops"
)

func main() {
	var commandValues int

	fmt.Println("*********************************")
	fmt.Println("Enter one of next commands to do the action: ")
	fmt.Println("1 -> Calculate factorial of provided number")
	fmt.Println("2 -> Reverse numbers order of provided number")
	fmt.Println("*********************************")

	fmt.Scan(&commandValues)

	switch commandValues {
	case 1:
		res := loops.Calculate()
		fmt.Println(res)
		break
	case 2:
		res := loops.Revers()
		fmt.Println(res)
		break
	default:
		fmt.Println("Invalid value provided")
	}

}
