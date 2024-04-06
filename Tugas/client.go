package main

import(
	"net"
	"fmt"
)

//client -> dial

func main(){
	dial, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = dial.Write([]byte("Hello Server!"))

	buffer := make([]byte,1024)

	n,err := dial.Read(buffer)

	fmt.Println(string(buffer[:n]))

	defer dial.Close()
}