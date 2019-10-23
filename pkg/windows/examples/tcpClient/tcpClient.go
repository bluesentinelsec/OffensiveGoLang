package main

import (
	"fmt"
	"log"

	"github.com/bluesentinelsec/OffensiveGoLang/pkg/windows/c2"
)

func main() {
	// specify remote IP Address
	ipAddr := "127.0.0.1:8080"

	// Initialize data to send to server
	data := []byte("Hello from Client!")

	// Connect to remote server
	conn, err := c2.TCPConnect(ipAddr)
	if err != nil {
		log.Fatal(err)
	}
	// close connection after function completes
	defer conn.Close()

	// write data to remote server
	err = c2.TCPWrite(conn, data)
	if err != nil {
		log.Fatal(err)
	}

	// read server response
	response, err := c2.TCPRead(conn)
	if err != nil {
		log.Fatal(err)
	}

	// print server response to standard out
	fmt.Println(string(response))

}
