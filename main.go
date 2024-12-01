package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error opening a connection to the star-db server")
		log.Fatal(err)
	}
	defer conn.Close()

	// Send a request
	request := []byte("This is a request from the star-db client\n")
	requestBytes, err := conn.Write(request)
	if err != nil {
		fmt.Println("Error sending a request to the star-db server: ", request)
		log.Fatal(err)
	}
	fmt.Printf("Sent %d bytes to the star-db server", requestBytes)

	// Read the response
	buf := make([]byte, 1024)
	responseBytes, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading the response from the star-db server: ", err)
		log.Fatal(err)
	}

	fmt.Println("Response from star-db server: ", string(buf[:responseBytes]))
}
