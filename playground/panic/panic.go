package main

import "fmt"

func main(){
	defer func(){
		if e:= recover(); e != nil {
			fmt.Println(e)
			fmt.Println("end")
		}
	}()
	var n int = 0
	fmt.Println(1/n)
}
