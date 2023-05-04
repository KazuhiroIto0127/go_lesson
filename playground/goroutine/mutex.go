package main

import (
	"fmt"
	"sync"
)

// mutexによるlockがないと2000にならないときが発生する
func main(){
	fmt.Println("test")
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func(){
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(n)
}
