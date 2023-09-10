package main

import (
	"fmt"
	"sync"
	"time"
)

func fib(n int) int{
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	start := time.Now() // 実行前の時刻を取得

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- fib(40)
		}()
	}

	go func(){
		for result := range ch {
			fmt.Println(result)
		}
	}()

	wg.Wait()
	close(ch)

	duration := time.Since(start) // 実行時間を計測
	fmt.Printf("Execution time: %v\n", duration)
}
