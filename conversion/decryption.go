package conversion

func Decrypt(s string, shift int) string {
	var decryptedText string

	for _, ch := range s {
		decryptedText += string(ch - rune(shift))
	}

	return decryptedText
}
