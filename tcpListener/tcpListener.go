package tcpListener

import (
	"fmt"
	"net"
)

func Accept(port int, connChan chan<- net.Conn) {

	fmt.Println("Listening on port ", port)

	addr := net.TCPAddr{}
	addr.Port = port
	tcpListener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println("Failed to set up Tcp listener!")
		return
	}

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("Failed accept connection!")
			return

		}

		connChan <- tcpConn
	}
}
