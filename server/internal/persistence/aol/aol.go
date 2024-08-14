package aol

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"io"
	"os"
	"strconv"
	"strings"
)

func Write(command string, key string, args []int64) error {
	file, err := os.OpenFile("resources/aol.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close file: %s because: %s", file.Name(), err.Error()))
		}
	}(file)

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(fmt.Sprintf("%s;%s;%v\n", command, key, args))

	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func Read() ([]command.Command, error) {
	var cmds []command.Command

	file, err := os.Open("resources/aol.bin")

	if err != nil {
		return []command.Command{}, errors.New("aol file not found. Skipping replay")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close file: %s because: %s", file.Name(), err.Error()))
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		argsStr := parts[2]
		args := parseArgs(argsStr)

		cmd, err := command.Create(parts[0], parts[1], args)
		if err != nil {
			return []command.Command{}, errors.New("aol file not found. Skipping replay")
		}

		cmds = append(cmds, cmd)
	}

	fmt.Println(cmds)

	return cmds, nil
}

func Write2(command string, key string, args []int64) error {
	length := byte(len(args))

	file, err := os.OpenFile("resources/aol.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close file: %s because: %s", file.Name(), err.Error()))
		}
	}(file)

	writer := bufio.NewWriter(file)

	err = binary.Write(writer, binary.LittleEndian, length)
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(command)))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(command))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(key)))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(key))
	if err != nil {
		return err
	}

	for _, arg := range args {
		err = binary.Write(writer, binary.LittleEndian, arg)

		if err != nil {
			return err
		}
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}

func Read2() ([]command.Command, error) {
	var cmds []command.Command

	file, err := os.Open("resources/aol.bin")

	if err != nil {
		return []command.Command{}, errors.New("aol file not found. Skipping replay")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close file: %s because: %s", file.Name(), err.Error()))
		}
	}(file)

	for {
		var length byte
		err = binary.Read(file, binary.LittleEndian, &length)

		if err != nil {
			if err == io.EOF {
				break
			}

			return []command.Command{}, errors.New("aol cannot be loaded. Skipping reply")
		}

		var commandLength byte
		err = binary.Read(file, binary.LittleEndian, &commandLength)

		if err != nil {
			return []command.Command{}, errors.New("aol cannot be loaded. Skipping replay")
		}

		commandName := make([]byte, commandLength)
		err = binary.Read(file, binary.LittleEndian, &commandName)

		if err != nil {
			return []command.Command{}, errors.New("aol cannot be loaded. Skipping replay")
		}

		var keyLength byte
		err = binary.Read(file, binary.LittleEndian, &keyLength)

		if err != nil {
			return []command.Command{}, errors.New("aol cannot be loaded. Skipping replay")
		}

		key := make([]byte, keyLength)
		err = binary.Read(file, binary.LittleEndian, &key)

		if err != nil {
			return []command.Command{}, errors.New("aol cannot be loaded. Skipping replay")
		}

		args := make([]int64, length)
		err = binary.Read(file, binary.LittleEndian, &args)

		if err != nil {
			return []command.Command{}, errors.New("aol cannot be loaded. Skipping replay")
		}

		cmd, _ := command.Create(string(commandName), string(key), args)

		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

func Replay() (int, error) {
	cmds, err := Read()

	numberOfReplayedCommands := 0

	if err != nil {
		return 0, err
	}

	for _, cmd := range cmds {
		t, err := tree.Create(cmd.Key, cmd.Type)

		if err != nil {
			return numberOfReplayedCommands, err
		}

		switch cmd.Name {
		case enum.BSTADD:
			t.Add(cmd.Args[0], false)
			numberOfReplayedCommands++
		case enum.BSTREM:
			t.Remove(cmd.Args[0], false)
			numberOfReplayedCommands++
		}
	}

	return numberOfReplayedCommands, nil
}

func parseArgs(args string) []int64 {
	argsTrimmedLeft := strings.TrimLeft(args, "[")
	argsTrimmed := strings.TrimRight(argsTrimmedLeft, "]")
	argsFormatted := strings.Split(argsTrimmed, " ")
	argsInt := make([]int64, 0)

	for _, arg := range argsFormatted {
		argInt, _ := strconv.Atoi(arg)
		argsInt = append(argsInt, int64(argInt))
	}

	return argsInt
}
