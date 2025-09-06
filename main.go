package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read from connection: %v", err)
	}

	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()
}

func main() {

	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}

		go do(conn)
	}

}
