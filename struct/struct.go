package main

import "fmt"

func main() {
	type User struct {
		Name string
		Age int
	}

	var user User
	user.Name = "Bob"
	user.Age = 18

	fmt.Println(user)

	kenny := User {
		Name: "Kenny",
		Age: 18,
	}
	fmt.Println(kenny)
}
