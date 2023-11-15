package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix  = "Bonjour, "
const spanish            = "Spanish"
const french             = "French"

func Hello(name, language string) string {

    if name == "" {
        return englishHelloPrefix + "World"
    }

    prefix := greetingPrefix(language)

    return prefix + name
}

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

func main() {
    fmt.Println(Hello("World", "English"))
}


