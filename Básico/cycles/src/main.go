package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i < 10; i++ {
		fmt.Println("For", i)
	}

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println("for while ", counter)
		counter++
	}

	// For backwards
	for i := 10; i > 0; i-- {
		fmt.Println("For Backwards", i)
	}

	// For Forever with a condition to break it
	counterForever := 0
	for {
		fmt.Println("Forever ", counterForever)
		counterForever++
		if counterForever == 5 {
			break
		}
	}

}
