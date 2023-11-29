package main

import (
	"fmt"

	. "github.com/aarbanas/go-learning/conversion"
	. "github.com/aarbanas/go-learning/errors"
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
	fmt.Println("5 -> Encryption")

	fmt.Println("*********************************")

	_, err := fmt.Scan(&commandValues)
	if err != nil {
		ScanErrorMessage(err)
		return
	}

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
	case 5:
		var s string
		shift := 5
		fmt.Println("Enter text you wish to encrypt")
		_, err := fmt.Scan(&s)
		if err != nil {
			ScanErrorMessage(err)
			return
		}

		encryptedText := Encrypt(s, shift)
		fmt.Println("Encrypted text: ", encryptedText)

		decryptedText := Decrypt(encryptedText, shift)
		fmt.Println("Decrypted text: ", decryptedText)

		break
	default:
		fmt.Println("Invalid value provided")
	}

}
