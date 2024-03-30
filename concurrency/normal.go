package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	start := time.Now() // 実行前の時刻を取得

	for i := 0; i < 5; i++ {
		result := fib(40)
		fmt.Printf("fib(40) = %d\n", result)
	}

	duration := time.Since(start) // 実行時間を計測
	fmt.Printf("Total execution time: %v\n", duration)
}
