package main

import (
	"io"
	"net"
	"fmt"
)

func main() {

	li, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
	}
	defer li.Close()

	conn, err := li.Accept()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(conn, "\nWhat are you doing ?\n")	
	conn.Close()

}
