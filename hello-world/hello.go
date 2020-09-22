package main

import "fmt"

// using const
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns greeting.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
