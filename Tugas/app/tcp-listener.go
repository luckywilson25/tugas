package main

import (
	"fmt"
	"net"

	"main.mod/handler"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	handler.ErrorHandler(err)

	fmt.Printf("Bound to %q\n", listener.Addr())
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handler.ErrorHandler(err)

		go dialHandler(conn)
	}
}

func dialHandler(conn net.Conn) {
	defer conn.Close()

	fmt.Println("one connection recived")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	handler.ErrorHandler(err)

	fmt.Println("Client: ", string(buffer[:n]))

	_, err = conn.Write([]byte("hello dari server udah keterima"))
	handler.ErrorHandler(err)
}
