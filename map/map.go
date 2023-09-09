package main

import "fmt"

func printAllMapValue(m map[string]int) {
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func main() {
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["orange"] = 3
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(len(m))
	delete(m, "apple")
	fmt.Println(m)

	printAllMapValue(m)
}
