package main

import "fmt"

const englishHelloPrefix = "Hello,"

func Hello(s string) string {
	if s == "" {
		return fmt.Sprintf("%s World", englishHelloPrefix)
	}
	return fmt.Sprintf("%s %s", englishHelloPrefix, s)
}

func main() {}
