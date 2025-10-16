package main

import "fmt"

func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Greet("World"))
}
