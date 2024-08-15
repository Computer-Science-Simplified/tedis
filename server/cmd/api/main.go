package main

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	persistencelistener "github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol/listeners"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/rdb"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
	storelistener "github.com/Computer-Science-Simplified/tedis/server/internal/store/listeners"
	"os"

	"net"

	"github.com/gookit/event"
)

func main() {
	fmt.Println("Starting Tedis...")

	restoreDatabase()

	capacity := store.Len()
	if capacity == 0 {
		capacity = 10
	}

	lru := store.NewLRU(capacity)

	for _, key := range store.Keys() {
		lru.Put(key)
	}

	addEventListeners(lru)

	listener, err := net.Listen("tcp", ":2222")

	if err != nil {
		panic(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close listener: %s. exiting", err.Error()))
			os.Exit(1)
		}
	}(listener)

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

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close connection: %s", err.Error()))
		}
	}(conn)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading from connection")
			return
		}

		commandName := string(buffer[:n])

		fmt.Printf("Received: %s", commandName)

		cmd, err := command.Parse(commandName)

		if err != nil {
			_, _ = conn.Write([]byte(err.Error() + "\n"))
		} else {
			result, err := cmd.Execute()

			if err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				return
			}

			_, err = conn.Write([]byte(result + "\n"))

			if err != nil {
				fmt.Println("Error writing to connection")
				return
			}
		}
	}
}

func addEventListeners(lru *store.LRU) {
	event.On("write_command_executed", event.ListenerFunc(func(e event.Event) error {
		store.CurrentUnsavedWriteCommands++

		data := e.Data()

		err := persistencelistener.AppendToAol(data)

		if err != nil {
			fmt.Println(fmt.Errorf("could not append to AOL: %s", err.Error()))
		}

		storelistener.EvictOldKeys(data, lru)

		return nil
	}))
}

func restoreDatabase() {
	fmt.Println("Checking environment for persistence layer")

	// Possible values: aol, rdb
	persistenceLayer, exists := os.LookupEnv("PERSISTENCE_LAYER")

	if !exists {
		fmt.Println("PERSISTENCE_LAYER not set. Defaulting to AOL")
		persistenceLayer = "aol"
	} else {
		fmt.Printf("PERSISTENCE_LAYER is set to %s\n", persistenceLayer)
	}

	if persistenceLayer == "aol" {
		fmt.Println("Replaying AOL...")

		count, err := aol.Replay()

		if err != nil {
			fmt.Println(fmt.Errorf("could not replay AOL: %s", err.Error()))
		}

		fmt.Printf("Replayed %d commands\n", count)
	}

	if persistenceLayer == "rdb" {
		fmt.Println("Reloading RDB...")

		count, err := rdb.Reload()

		if err != nil {
			fmt.Println(fmt.Errorf("could not reload RDB: %s", err.Error()))
		}

		fmt.Printf("Reloaded %d items\n", count)
	}
}
