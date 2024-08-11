package main

import (
	"fmt"
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/listeners"
	"mmartinjoo/trees/store"

	"mmartinjoo/trees/command"
	"net"

	"github.com/gookit/event"
)

func main() {
	fmt.Println("Starting Tedis...")

	fmt.Println("Replaying AOL...")
	aol.Replay()
	fmt.Println("DONE")

	// fmt.Println("Reloading RDB...")
	// rdb.Reload()
	// fmt.Println("DONE")

	lru := store.NewLRU(len(store.Store))

	for key := range store.Store {
		lru.Put(key)
	}

	event.On("write_command_executed", event.ListenerFunc(func(e event.Event) error {
		store.CurrentUnsavedWriteCommands++

		data := e.Data()

		listeners.LogWriteCommand(data)

		return nil
	}))

	event.On("command_executed", event.ListenerFunc(func(e event.Event) error {
		data := e.Data()

		listeners.EvictOldKeys(data, lru)

		return nil
	}))

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

		command, err := command.Parse(commandName)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		}

		result := command.Execute()

		_, err = conn.Write([]byte(result + "\n"))

		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
	}
}
