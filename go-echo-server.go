package main

import (
	"fmt"
	"github.com/meros/go-tcplistener"
	"net"
)

func Serve(conn net.Conn) {
	defer conn.Close()

	for {
		var buffer [2048]byte
		n, err := conn.Read(buffer[:])

		if err != nil {
			fmt.Println("Failed to read from connection...")
			return
		}

		if n == 0 {
			fmt.Println("Connection closed...")
			return
		}

		fmt.Println("Got ", n, " bytes on connection")

		conn.Write(buffer[0 : n-1])
	}
}

func handleClients(connChan <-chan net.Conn) {
	var tcpConn net.Conn
	for {
		tcpConn = <-connChan

		fmt.Println("Accepted new client: ", tcpConn.RemoteAddr().String())

		go Serve(tcpConn)
	}
}

func main() {
	connChan := make(chan net.Conn)
	go handleClients(connChan)

	err := tcplistener.Accept(8080, connChan)
	if err != nil {
		fmt.Println("Failed to start tcplistener:", err)
	}
}
