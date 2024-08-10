package aol

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"mmartinjoo/trees/command"
	"os"
)

/*
 * SET key1 value1
 *
  3
  3
  SET
  4
  key1
  6
  value1
*/

func Write(command string, key string, args []int64) {
	var length byte = byte(len(args) + 2)

	file, err := os.OpenFile("aol.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	err = binary.Write(writer, binary.LittleEndian, length)
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(command)))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(command))	
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(key)))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(key))	
	if err != nil {
		panic(err)
	}

	for _, arg := range args {
		err = binary.Write(writer, binary.LittleEndian, arg)

		if err != nil {
			panic(err)
		}
	}

	err = writer.Flush()

    if err != nil {
        panic(err)
    }
}

func Read() ([]command.Command, error) {
	var commands []command.Command

	file, err := os.Open("aol.bin")

	if err != nil {
		return []command.Command{}, errors.New("AOL file not found. Skipping replay.")
	}

	defer file.Close()

	for {
		var length byte

		err = binary.Read(file, binary.LittleEndian, &length)

		if err != nil {
			if err == io.EOF {
				break
			}

			return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
		}

		var commandLength byte

		err = binary.Read(file, binary.LittleEndian, &commandLength)

		if err != nil {
			return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
		}

		var commandName string
		for i := 0; i < int(commandLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
			}

			commandName += string(c)
		}

		var keyLength byte

		err = binary.Read(file, binary.LittleEndian, &keyLength)

		if err != nil {
			return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
		}

		var key string
		for i := 0; i < int(keyLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
			}

			key += string(c)
		}

		var args []int64

		for i := 0; i < int(length) - 2; i++ {
			var arg int64
			err = binary.Read(file, binary.LittleEndian, &arg)

			if err != nil {
				return []command.Command{}, errors.New("AOL cannot be loaded. Skipping replay.")
			}

			args = append(args, arg)
		}

		cmd, _ := command.Create(commandName, key, args)

		commands = append(commands, cmd)
	}

	return commands, nil
}
