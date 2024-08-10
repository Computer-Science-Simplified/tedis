package commandparser

import (
	"errors"
	"fmt"
	"mmartinjoo/trees/command"
	"strconv"
	"strings"
)

func Parse(line string) (command.Command, error) {
	trimmedCommand := strings.TrimRight(line, "\n")

	parts := strings.Split(trimmedCommand, " ")

	if len(parts) < 2 {
		return command.Command{}, errors.New("command and key are required")
	}

	name := parts[0]

	key := parts[1]

	args := parts[2:]

	if name == "BTADD" && len(args) != 1 {
		return command.Command{}, fmt.Errorf("btadd requires exactly 1 argument but %d given", len(args))
	}

	if name == "BTEXISTS" && len(args) != 1 {
		return command.Command{}, fmt.Errorf("btexists requires exactly 1 argument but %d given", len(args))
	}

	if name == "BTGETALL" && len(args) != 0 {
		return command.Command{}, fmt.Errorf("btexists requires exactly 0 argument but %d given", len(args))
	}

	formattedArgs := make([]int64, 0)

	if (len(args) > 0) {
		formattedArgs = make([]int64, len(args) - 1)

		for _, arg := range args {
			val, _ := strconv.Atoi(arg)
	
			formattedArgs = append(formattedArgs, int64(val))
		}
	}

	cmd, err := command.Create(name, key, formattedArgs)

	if err != nil {
		return command.Command{}, err
	}

	return cmd, nil
}