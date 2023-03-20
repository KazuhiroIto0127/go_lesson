package main

import "fmt"

func plus(a int, b int) int{
	return a+b
}

func multiReturn() (int, int) {
	return 3, 7;
}

func main() {
	fmt.Println(plus(1, 2))
	a, b := multiReturn()
	fmt.Println(a, b)
}
