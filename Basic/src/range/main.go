package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("is palindrom")
	} else {
		fmt.Println("is not a palindrom")
	}
}
func main() {
	slice := []string{"hello", "how", "are", "you"}
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// palindroms

	isPalindrome("Ama")

}
