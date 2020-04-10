package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "Hello! How are you ?")
}
