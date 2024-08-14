package aol

import (
	"bufio"
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"os"
	"strconv"
	"strings"
	"time"
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
		return []command.Command{}, err
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

		cmd, err := command.Create(strings.ToUpper(parts[0]), parts[1], args)
		if err != nil {
			return []command.Command{}, err
		}

		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

func Replay() (int, error) {
	start := time.Now().Unix()

	cmds, err := Read()

	numberOfReplayedCommands := 0

	if err != nil {
		return 0, err
	}

	for _, cmd := range cmds {
		t, err := tree.Create(cmd.Key, cmd.Type)

		if err != nil {
			fmt.Printf("unable to create tree from command: %s/n", cmd)
			continue
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

	end := time.Now().Unix()

	fmt.Printf("%d\n", end-start)

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
