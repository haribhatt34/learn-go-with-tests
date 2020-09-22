package main

import "fmt"

// using const
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

/*
 * we can use a named variable with return type, if we want to be more verbal.
 * like (prefix string), it will create a variable prefix in greetingPrefix
 * and assign "zero" value. "" and 0 for string and int respectively.
 * we can return without any variable in the end, as prefix will be returned.
 *
 * func with lowercase are private, uppercase as public.
 * since greetingPrefix is a helper, we are keeping it private.
 * while Hello is used by external methods like testing module.
 */

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello returns greeting.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
