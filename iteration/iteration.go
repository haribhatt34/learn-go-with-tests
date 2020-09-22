package iteration

// Repeat returns character repeated 5 times
func Repeat(character string) (repeatedString string) {
	for i := 0; i < 5; i++ {
		repeatedString += character
	}
	return
}
