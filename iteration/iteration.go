package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times
func Repeat(character string) (repeatedString string) {
	for i := 0; i < repeatCount; i++ {
		repeatedString += character
	}
	return
}
