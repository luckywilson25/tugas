package main

import(
	"net"
	"fmt"
)

func main() {
	//server -> listen
	//client -> dial

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for{
		clientConn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleClient(clientConn)
	}
}

func handleClient(clientConn net.Conn){
	defer clientConn.Close()

	//fmt.Println("sudah tersambung")

	buffer := make([]byte,1024)

	n, err := clientConn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string (buffer[:n]))

	//cara respon server ke client
	_, err = clientConn.Write([]byte("hello dari server! udh ke terima"))

}