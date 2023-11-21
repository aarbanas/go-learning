package conversion

import "fmt"

func Encrypt() {
	var s string
	var shift int
	var encryptedText string

	fmt.Println("First type text than enter the shift value")
	fmt.Scan(&s, &shift)

	for _, ch := range s {
		encryptedText += string(ch + rune(shift))
	}
	fmt.Print(encryptedText)
}
