package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(wg *sync.WaitGroup, i int, msg string){
	defer wg.Done()

	fmt.Println("start", msg, i)
	time.Sleep(time.Second)
	fmt.Println("end", msg, i)
}

func main(){
	var wg sync.WaitGroup

	message := "hi"

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sendMessage(&wg, i, message)
	}

	message = "ho"
	fmt.Println(message)

	wg.Wait()
}
