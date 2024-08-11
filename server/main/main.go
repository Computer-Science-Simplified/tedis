package main

import (
	"fmt"
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/cache"
	"mmartinjoo/trees/store"

	commandparser "mmartinjoo/trees/command_parser"
	"mmartinjoo/trees/rdb"
	"net"

	"github.com/gookit/event"
)

func main() {
	fmt.Println("Starting Tedis...")

	fmt.Println("Replaying AOL...")
	aol.Replay()
	fmt.Println("DONE")

	var capacity int

	if len(store.Store) == 0 {
		capacity = 6
	} else {
		capacity = min(len(store.Store), 6)
	}

	lru := cache.NewLRU(capacity)

	for key := range store.Store {
		lru.Put(key)
	}

	numberOfWriteCommands := 0

	event.On("write_command_executed", event.ListenerFunc(func(e event.Event) error {
		numberOfWriteCommands++

		data := e.Data()
		command, _ := data["command"].(string)
		key, _ := data["key"].(string)
		args, _ := data["args"].([]int64)

		aol.Write(command, key, args)

		if numberOfWriteCommands%10 == 0 {
			rdb.Persist()
		}

		return nil
	}))

	event.On("command_executed", event.ListenerFunc(func(e event.Event) error {
		data := e.Data()
		key, _ := data["key"].(string)

		_, err := lru.Get(key)

		if err != nil {
			lru.Put(key)
		}

		if len(store.Store) >= 5 {
			evictionTargets := lru.GetLeastRecentlyUsed(len(store.Store) - 5)

			for _, evictionTarget := range evictionTargets {
				delete(store.Store, evictionTarget)
				lru.Remove(evictionTarget)
			}

		}

		return nil
	}))

	event.On("key_created", event.ListenerFunc(func(e event.Event) error {
		// data := e.Data()
		// key, _ := data["key"].(string)

		// TODO: same as command_execute

		fmt.Println(len(store.Store))

		return nil
	}))

	fmt.Println(lru.Map)
	fmt.Println(lru.Items)

	// fmt.Println("Reloading RDB...")
	// rdb.Reload()
	// fmt.Println("DONE")

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

		_, err = conn.Write([]byte(result + "\n"))

		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
	}
}
