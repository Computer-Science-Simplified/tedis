package command

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"strings"
)

func Create(name string, key string, args []int64) (Command, error) {
	if strings.HasPrefix(name, "BST") {
		return Command{
			Name: name,
			Key:  key,
			Args: args,
			Type: enum.BinarySearchTree,
		}, nil
	}

	return Command{}, fmt.Errorf("command not found %s", name)
}
