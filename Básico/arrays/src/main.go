package main

import "fmt"

func main() {
	// declare an array
	var array [4]int
	fmt.Println(array)
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	// slice
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:3])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 6)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{7, 8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
