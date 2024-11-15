package aol

import (
	"bufio"
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const fileName = "./resources/aol.log"

func Append(cmd command.Command) error {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("unable to close file: %s because: %s", file.Name(), err.Error()))
		}
	}(file)

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(
		fmt.Sprintf("%s;%s;%s\n", cmd.GetParams().Name, cmd.GetParams().Key, convertArgs(cmd.GetParams().Args)),
	)

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

	file, err := os.Open(fileName)

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
	start := time.Now().UnixMilli()

	cmds, err := Read()

	numberOfReplayedCommands := 0

	if err != nil {
		return 0, err
	}

	for _, cmd := range cmds {
		_, err := cmd.Execute(false)
		if err != nil {
			fmt.Printf(
				"unable to create tree from command: %s\n",
				cmd.String(),
			)

			continue
		}

		numberOfReplayedCommands++
	}

	end := time.Now().UnixMilli()

	fmt.Printf("AOL replayed in %dms\n", end-start)

	return numberOfReplayedCommands, nil
}

func parseArgs(args string) []int64 {
	argsFormatted := strings.Split(args, ",")
	argsInt := make([]int64, len(argsFormatted))

	for i, arg := range argsFormatted {
		argInt, _ := strconv.Atoi(arg)
		argsInt[i] = int64(argInt)
	}

	return argsInt
}

func convertArgs(args []int64) string {
	strArgs := make([]string, len(args))
	for i, v := range args {
		strArgs[i] = fmt.Sprint(v)
	}

	return strings.Join(strArgs, ",")
}
