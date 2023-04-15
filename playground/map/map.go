package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(len(m))
	delete(m, "apple")
	fmt.Println(m)
}
