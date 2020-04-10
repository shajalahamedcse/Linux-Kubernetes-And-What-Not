package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(data)
	}
	fmt.Println(string(data))
}
