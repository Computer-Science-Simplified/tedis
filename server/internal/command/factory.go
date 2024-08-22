package command

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command/bst"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"strings"
)

func Create(name string, key string, args []int64) (Command, error) {
	if strings.HasPrefix(name, "BST") {
		cp := &types.CommandParams{
			Name: name,
			Key:  key,
			Args: args,
			Type: enum.BinarySearchTree,
		}

		switch name {
		case enum.BSTADD:
			return bst.BSTAdd{Params: cp}, nil
		case enum.BSTGETALL:
			return bst.BSTGetAll{Params: cp}, nil
		case enum.BSTREM:
			return bst.BSTRem{Params: cp}, nil
		case enum.BSTEXISTS:
			return bst.BSTExists{Params: cp}, nil
		}
	}

	return nil, fmt.Errorf("command not found %s", name)
}
