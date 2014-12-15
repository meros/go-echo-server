package echoServer

import (
	"fmt"
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
