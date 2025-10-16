package main

import "fmt"

func Greet(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "French":
		return "Bonjour, " + name
	case "Spanish":
		return "Hola, " + name
	default:
		return "Hello, " + name
	}
}

func main() {
	fmt.Println(Greet("World", "English"))
}
