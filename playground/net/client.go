package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connectin:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "hello server! \n")
}
