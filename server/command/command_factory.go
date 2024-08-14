package command

import (
	"fmt"
	"mmartinjoo/trees/internal/trees"
	"strings"
)

func Create(name string, key string, args []int64) (Command, error) {
	if strings.HasPrefix(name, "BST") {
		return Command{
			Name: name,
			Key:  key,
			Args: args,
			Type: trees.BinarySearchTree,
		}, nil
	}

	return Command{}, fmt.Errorf("command not found %s", name)
}
