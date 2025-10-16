package main

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Japanese":
		return "こんにちは、" + name
	default:
		return "Hello, " + name
	}
}

func main() {
	println(Hello("World", ""))
}
