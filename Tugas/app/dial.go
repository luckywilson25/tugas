package main

import (
	"fmt"
	"net"
	"time"

	"main.mod/handler"
)

func main() {
	dial, err := net.DialTimeout("tcp", "localhost:1234", time.Second*5)
	handler.ErrorHandler(err)
	defer dial.Close()

	dial.SetDeadline(time.Now().Add(time.Second * 5))

	netErr, ok := err.(net.Error)

	_, err = dial.Write([]byte("hello server"))
	handler.ErrorHandler(err)

	buffer := make([]byte, 1024)
	n, err := dial.Read(buffer)

	fmt.Println("server: ", string(buffer[:n]))

	if err != nil {
		if ok && netErr.Timeout() {
			fmt.Println("connection timeout")
			return
		} else {
			handler.ErrorHandler(err)
		}
	}
}
