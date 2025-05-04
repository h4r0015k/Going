package main

import "fmt"

const engPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getGreeting(lang) + name
}

func getGreeting(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = engPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Nikhil", ""))
}
