package main

import (
	"./echoServer"
	"./tcpListener"
	"fmt"
	"net"
)

func handleClients(connChan <-chan net.Conn) {
	var tcpConn net.Conn
	for {
		tcpConn = <-connChan

		fmt.Println("Accepted new client: ", tcpConn.RemoteAddr().String())

		go echoServer.Serve(tcpConn)
	}
}

func main() {
	connChan := make(chan net.Conn)
	go handleClients(connChan)

	tcpListener.Accept(8080, connChan)
}
