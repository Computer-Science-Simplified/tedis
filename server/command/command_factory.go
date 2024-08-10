package command

import (
	"fmt"
	"strings"
)

func Create(name string, key string, args []int64) (Command, error) {
	if strings.HasPrefix(name, "BT") {
		return Command{
			Name: name,
			Key: key,
			Args: args,
			Type: "binary_tree",
		}, nil
	}

	return Command{}, fmt.Errorf("command not found %s", name)
}