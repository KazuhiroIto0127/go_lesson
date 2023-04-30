package main

import (
	"fmt"
	"log"
	"os"
)

func ConfirmDeferOrder(){
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
