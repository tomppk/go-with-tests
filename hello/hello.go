package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "Ni hao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// A named return value (prefix string) for function. We can return whatever
// is set to prefix by just calling return rather than return prefix
// Creates a variable called prefix in function and by default it will be assigned
// zero value for int 0 and string ""
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
		return
}



func main() {
	fmt.Println(Hello("Chris", ""))
}
