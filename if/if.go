package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random int = rand.Intn(10)

	if random >= 7 {
		fmt.Println("7以上：", random)
	} else if random >= 3 {
		fmt.Println("3以上６以下：", random)
	} else {
		fmt.Println("3以下：", random)
	}
}
