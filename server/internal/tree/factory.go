package tree

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/model"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func Create(key string, treeType string) (model.Tree, error) {
	if item, ok := store.Get(key); ok {
		return item, nil
	}

	if treeType == enum.BinarySearchTree {
		tree := &BST{
			Key: key,
		}

		store.Set(key, tree)

		return tree, nil
	}

	return nil, fmt.Errorf("tree type not found: %s", treeType)
}
