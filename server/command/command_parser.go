package command

import (
	"errors"
	"fmt"
	"mmartinjoo/trees/commands"
	"strconv"
	"strings"
)

func Parse(line string) (Command, error) {
	trimmedCommand := strings.TrimRight(line, "\n")

	parts := strings.Split(trimmedCommand, " ")

	if len(parts) < 2 {
		return Command{}, errors.New("command and key are required")
	}

	name := parts[0]

	key := parts[1]

	args := parts[2:]

	if name == commands.BTADD && len(args) != 1 {
		return Command{}, fmt.Errorf("%s requires exactly 1 argument but %d given", commands.BTADD, len(args))
	}

	if name == commands.BTEXISTS && len(args) != 1 {
		return Command{}, fmt.Errorf("%s requires exactly 1 argument but %d given", commands.BTEXISTS, len(args))
	}

	if name == commands.BTGETALL && len(args) != 0 {
		return Command{}, fmt.Errorf("%s requires exactly 0 argument but %d given", commands.BTGETALL, len(args))
	}

	if name == commands.BTREM && len(args) != 1 {
		return Command{}, fmt.Errorf("%s requires exactly 1 argument but %d given", commands.BTREM, len(args))
	}

	formattedArgs := make([]int64, 0)

	if (len(args) > 0) {
		formattedArgs = make([]int64, len(args) - 1)

		for _, arg := range args {
			val, _ := strconv.Atoi(arg)
	
			formattedArgs = append(formattedArgs, int64(val))
		}
	}

	cmd, err := Create(name, key, formattedArgs)

	if err != nil {
		return Command{}, err
	}

	return cmd, nil
}