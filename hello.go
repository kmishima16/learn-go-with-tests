package main

const japanesePrefix = "こんにちは、"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	
	prefix := getGreetingPrefix(language)
	return prefix + name
}

func getGreetingPrefix(language string) string {
	switch language {
	case "Japanese":
		return japanesePrefix
	case "Spanish":
		return spanishPrefix
	default:
		return englishPrefix
	}
}

func main() {
	println(Hello("World", ""))
}
