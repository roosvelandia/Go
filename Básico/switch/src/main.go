package main

import "fmt"

func main() {

	switch module := 5 % 2; module {
	case 0:
		fmt.Println("is odd")
	default:
		fmt.Println("is even")
	}

	// without condition
	value := -1
	switch {
	case value > 100:
		fmt.Println("is greater than 100")
	case value < 0:
		fmt.Println("is less than 0")
	default:
		fmt.Println("is between 0 and 100")
	}
}
