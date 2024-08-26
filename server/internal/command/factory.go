package command

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command/binarytree"
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
			return bst.NewBSTAdd(cp), nil
		case enum.BSTGETALL:
			return bst.NewBSTGetAll(cp), nil
		case enum.BSTREM:
			return bst.NewBSTRem(cp), nil
		case enum.BSTEXISTS:
			return bst.NewBSTExists(cp), nil
		}
	}

	if strings.HasPrefix(name, "BT") {
		cp := &types.CommandParams{
			Name: name,
			Key:  key,
			Args: args,
			Type: enum.BinaryTree,
		}

		switch name {
		case enum.BTADD:
			return binarytree.NewBTAdd(cp), nil
		}
	}

	return nil, fmt.Errorf("command not found %s", name)
}
