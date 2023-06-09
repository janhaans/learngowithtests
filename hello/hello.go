package hello

import "fmt"

const (
	englishHelloPrefix = "Hello,"
	spanishHelloPrefix = "Hola,"
	frenchHelloPrefix  = "Bonjour,"
)

func Hello(person, language string) string {
	prefix := GetPrefix(language)

	if person == "" {
		return fmt.Sprintf("%s World", prefix)
	}
	return fmt.Sprintf("%s %s", prefix, person)
}

func GetPrefix(language string) (prefix string) {
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
