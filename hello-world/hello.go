package main

import "fmt"

// using const
const englishHelloPrefix = "Hello, "

// Hello returns greeting.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
