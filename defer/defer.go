package main

import (
	"fmt"
	"log"
	"os"
)

func ConfirmDeferOrder(){
	var n = 100
	defer fmt.Println("普通の関数", n)
	defer func(){
		fmt.Println("無名関数", n)
	}()
	n = 200
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func ReadFileAndPrint(){
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // あとでよばれる

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
}

func main(){
	ConfirmDeferOrder()

	ReadFileAndPrint()
}
