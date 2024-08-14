package command

import (
	"fmt"
	"mmartinjoo/trees/internal/tree"
	"strings"
)

func Create(name string, key string, args []int64) (Command, error) {
	if strings.HasPrefix(name, "BST") {
		return Command{
			Name: name,
			Key:  key,
			Args: args,
			Type: tree.BinarySearchTree,
		}, nil
	}

	return Command{}, fmt.Errorf("command not found %s", name)
}
