package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Arguemnts Invalid")
		fmt.Println("Format is <script> <address> <port>")
		fmt.Println("Example: ./tcp-client 127.0.0.1 8080")
		return
	}

	address := os.Args[1]
	port := os.Args[2]
	startServer(address, port)
}

func startServer(address string, port string) {
	conn, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status, err)
}
