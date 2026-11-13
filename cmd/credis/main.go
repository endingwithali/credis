package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello, Go")

	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("Error accepting incoming connection: %+v\n", err)
		}
		defer conn.Close()
		bytes, err := io.ReadAll(conn)
		if err != nil {
			fmt.Errorf("Error reading incoming: %+v\n", err)
		}
		fmt.Printf("RECEIVED \n %s", bytes)
		_, err = conn.Write([]byte("RECEIVED"))
		if err != nil {
			fmt.Errorf("Error replying to incoming: %+v\n", err)
		}
		_, err = conn.Write(bytes)
		if err != nil {
			fmt.Errorf("Error echoing to incoming: %+v\n", err)
		}

	}
}
