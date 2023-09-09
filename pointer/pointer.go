package main

import "fmt"

func main() {
	var test int = 1
	var test_pointer *int = &test
	fmt.Println(test_pointer)

	*test_pointer = 2
	fmt.Println(test)
}
