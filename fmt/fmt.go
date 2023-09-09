package main

import "fmt"

type F struct{
	Name string
	Age int
}

func main(){
	f := &F {
		Name: "John",
		Age: 20,
	}

	fmt.Printf("%v\n", f)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%T\n", f)
}
