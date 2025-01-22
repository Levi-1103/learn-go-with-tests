package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		return "Hello, World!"
	}
	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Levi", ""))
}
