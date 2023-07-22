package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("Is 1")
	} else {
		fmt.Println("Isn't 1")
	}

	// With and
	if value1 == 1 && value2 == 2 {
		fmt.Println("Is true")
	}

	// With or
	if value1 == 0 || value2 == 2 {
		fmt.Println("fit at least one")
	}
	// Cast text to number
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)
	// Is even
	even := 2
	if even%2 == 0 {
		fmt.Println("is even")
	}
	// Is odd
	odd := 1
	if odd%2 != 0 {
		fmt.Println("is none")
	}
	// pwd and user
	pwd := "1234"
	user := "go"
	if user == "go" && pwd == "1234" {
		fmt.Println("Credentials correct")
	}
	// Cast text to number error
	valueError, err := strconv.Atoi("aaa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("error", valueError)

}
