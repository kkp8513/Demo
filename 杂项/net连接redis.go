package main

import (
	"fmt"
	"net"
	_"context"
	_"time"
	"log"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(str)
	}
}
