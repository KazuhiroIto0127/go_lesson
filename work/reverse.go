package main

import (
	"fmt"
	"strings"
)

// 与えられた文字列を反転させた結果を表示するプログラム
func main(){
	str := "abcdef"
	strArray := strings.Split(str,"")
	strLength := len(strArray)
	reversedStr := ""

	for i := 0; i < strLength; i++ {
		// fmt.Println(strLength-i-1)
		reversedStr = reversedStr + strArray[strLength-i-1]
	}

	fmt.Println(reversedStr)
}
