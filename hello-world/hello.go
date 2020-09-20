package main

import "fmt"

// using const
const englishHelloPrefix = "Hello, "

// Hello returns greeting.
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
