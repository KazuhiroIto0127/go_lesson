package main

import "fmt"

func main() {
	slice := []string{
		"a",
		"b",
		"c",
	}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	mapVal := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for key, value := range mapVal {
		fmt.Println(key, value)
	}
}
