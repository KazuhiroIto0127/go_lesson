package main

import (
	"fmt"
	"time"
)

func server1(ch chan string){
	time.Sleep(1 * time.Second)
	ch <- "server1 : hello"
}

func server2(ch chan string){
	time.Sleep(2 * time.Second)
	ch <- "server2 : hello"
}

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- ch1:
			fmt.Println(msg1)
		case msg2 := <- ch2:
			fmt.Println(msg2)
		}
	}
}
