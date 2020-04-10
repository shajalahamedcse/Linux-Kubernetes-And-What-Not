package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(conn)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said: %s\n", ln)
	}
	defer conn.Close()
}
