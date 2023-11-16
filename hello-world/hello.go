package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

//nolint:nonamedreturns
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

//nolint:forbidigo
func main() {
	fmt.Println(Hello("World", "English"))
}
