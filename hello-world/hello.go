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

    prefix := englishHelloPrefix

    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    }

    return prefix + name
}

func main() {
    fmt.Println(Hello("World", "English"))
}


