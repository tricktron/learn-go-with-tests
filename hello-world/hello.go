package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish            = "Spanish"

func Hello(name, language string) string {

    if name == "" {
        return englishHelloPrefix + "World"
    }

    if language == spanish {
        return spanishHelloPrefix + name
    }

    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("World", "English"))
}


