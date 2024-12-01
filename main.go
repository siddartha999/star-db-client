package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func handleConnection(conn *net.Conn, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Send a request
	request := []byte("This is a request from the star-db client\n")
	requestBytes, err := (*conn).Write(request)
	if err != nil {
		fmt.Println("Error sending a request to the star-db server: ", request)
		log.Fatal(err)
	}
	fmt.Printf("Sent %d bytes to the star-db server. index: %d \n", requestBytes, index)

	// Read the response
	buf := make([]byte, 1024)
	responseBytes, err := (*conn).Read(buf)
	if err != nil {
		fmt.Println("Error reading the response from the star-db server: ", err)
		log.Fatal(err)
	}
	fmt.Printf("index response: %d \n.", index)
	fmt.Println("Response from star-db server: \n", string(buf[:responseBytes]))
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error opening a connection to the star-db server")
		log.Fatal(err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go handleConnection(&conn, i, &wg)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
	conn.Close()
}
