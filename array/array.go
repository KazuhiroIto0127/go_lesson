package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println("長さ", len(array))
	fmt.Println(array[3])
}
