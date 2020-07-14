package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := os.Args[1]
	startServer(port)
}

func startServer(port string) {
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status, err)
}
