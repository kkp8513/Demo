package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	In, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := In.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Fprintf(conn, "ok")
	conn.Close()
}