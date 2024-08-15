package command

import (
	"errors"
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"strconv"
	"strings"
)

func Parse(line string) (*Command, error) {
	trimmedCommand := strings.TrimRight(line, "\n")

	parts := strings.Split(trimmedCommand, " ")

	if len(parts) < 2 {
		return nil, errors.New("command and key are required")
	}

	name := strings.ToUpper(parts[0])

	fmt.Println(name)

	key := parts[1]

	args := parts[2:]

	if name == enum.BSTADD && len(args) != 1 {
		return nil, fmt.Errorf("%s requires exactly 1 argument but %d given", enum.BSTADD, len(args))
	}

	if name == enum.BSTEXISTS && len(args) != 1 {
		return nil, fmt.Errorf("%s requires exactly 1 argument but %d given", enum.BSTEXISTS, len(args))
	}

	if name == enum.BSTGETALL && len(args) != 0 {
		return nil, fmt.Errorf("%s requires exactly 0 argument but %d given", enum.BSTGETALL, len(args))
	}

	if name == enum.BSTREM && len(args) != 1 {
		return nil, fmt.Errorf("%s requires exactly 1 argument but %d given", enum.BSTREM, len(args))
	}

	formattedArgs := make([]int64, 0)

	if len(args) > 0 {
		formattedArgs = make([]int64, len(args)-1)

		for _, arg := range args {
			val, _ := strconv.Atoi(arg)

			formattedArgs = append(formattedArgs, int64(val))
		}
	}

	cmd, err := Create(name, key, formattedArgs)

	if err != nil {
		return nil, err
	}

	return cmd, nil
}
