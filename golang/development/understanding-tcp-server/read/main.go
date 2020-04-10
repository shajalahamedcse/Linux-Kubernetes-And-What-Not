package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
)

func main(){
	li ,err := net.Listen("tcp",":8090")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()

	conn , err := li.Accept()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}