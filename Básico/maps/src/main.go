package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Joseph"] = 14
	m["Jhon"] = 20
	fmt.Println(m)

	// go throug
	for i, v := range m {
		fmt.Println(i, v)
	}

	// find value existent
	value := m["Joseph"]
	fmt.Println(value)
	// fin value non existent
	value = m["Jseph"]
	fmt.Println(value)
	// enhance my research when exists
	value, ok := m["Joseph"]
	fmt.Println(value, ok)
	// enhance my research when doesn't exist
	value, ok = m["Jseph"]
	fmt.Println(value, ok)
}
