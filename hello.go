package main

import "fmt"

const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
