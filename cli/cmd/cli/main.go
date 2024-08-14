package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2222")
	if err != nil {
		fmt.Println("Couldn't connect to server at localhost:2222")
		os.Exit(1)
	}

	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Couldn't read input")
			os.Exit(1)
		}

		_, err = conn.Write([]byte(command))

		if err != nil {
			fmt.Println("Couldn't send command")
			os.Exit(1)
		}

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Couldn't get response from server")
			os.Exit(1)
		}

		fmt.Printf("%s", string(buffer[:n]))
	}
}