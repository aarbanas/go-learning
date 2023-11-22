package conversion

func Encrypt(s string, shift int) string {
	var encryptedText string

	for _, ch := range s {
		encryptedText += string(ch + rune(shift))
	}

	return encryptedText
}
