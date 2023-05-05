package main

import (
	"fmt"
	"log"
	"time"
)

func main(){
	fmt.Println("test")
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05 JST"))
	fmt.Println(now.Format(time.RFC3339))

	var s = "2022/05/05 14:51:30"
	d, err := time.Parse("2006/01/02 15:04:05", s)
	if err != nil {
		log.Fatal("error parse error")
	}
	fmt.Println(d)

	duration, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(duration * 3)
}
