package main

import "fmt"

type User struct {
	name string
}

func (u User) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}

func main() {
	kazu := User{ name: "kazu" }
	fmt.Println(kazu.name, kazu.cal(70, 168))

	var v Value = 1
	fmt.Println(v.Add(2))
}
