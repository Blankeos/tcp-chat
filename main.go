package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()

	log.Printf("Started server on :8001")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Unable to start server: %s", err.Error())
			continue
		}
		c := s.newClient(conn)
		go c.readInput()
	}
}
