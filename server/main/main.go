package main

import (
	"fmt"
	"mmartinjoo/trees/aol"
	// aolreplayer "mmartinjoo/trees/aol_replayer"
	commandparser "mmartinjoo/trees/command_parser"
	"mmartinjoo/trees/rdb"
	"net"
	"github.com/gookit/event"
)

func main() {
	fmt.Println("Starting Tedis...")

	event.On("write_command_executed", event.ListenerFunc(func(e event.Event) error {
		data := e.Data()
		command, _ := data["command"].(string)
		key, _ := data["key"].(string)
		args, _ := data["args"].([]int64)

		aol.Write(command, key, args)

		return nil
	}));

	// fmt.Println("Replaying AOL...")
	// aolreplayer.Replay()
	// fmt.Println("DONE")

	fmt.Println("Reloading RDB...")
	rdb.Reload()
	fmt.Println("DONE")


	listener, err := net.Listen("tcp", ":2222")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Tedis is listening on port 2222")
	fmt.Println("Ready to accept connections")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Connection established with: %s\n", conn.RemoteAddr().String())

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading from connection")
			return
		}

		commandName := string(buffer[:n])

		fmt.Printf("Received: %s", commandName)

		command, err := commandparser.Parse(commandName)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		}

		result := command.Execute()

		rdb.Persist()

		_, err = conn.Write([]byte(result + "\n"))

		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
	}
}