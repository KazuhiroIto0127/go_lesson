package main

import "fmt"

func PrintDetail(v interface{}){
	switch t := v.(type){
	case int, int32, int64:
		fmt.Println("int/int32/int64:", t)
	case string:
		fmt.Println("string :", t)
	default:
		fmt.Println("知らない型")
	}
}

func main(){
	var i interface{}

	i = 1
	t := i.(int)
	fmt.Println(t)

	i = "こんにちは"
	fmt.Println(i)

	PrintDetail("test")
	PrintDetail(1111)

	// go ver18.1以降 interface{} = any でok

	var v any
	v = 1
	fmt.Println(v)
}
