package iteration

// Repeat returns character repeated 5 times
func Repeat(character string, repeatCount int) (repeatedString string) {
	for i := 0; i < repeatCount; i++ {
		repeatedString += character
	}
	return
}
