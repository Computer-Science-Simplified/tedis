package tree

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/model"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func Create(key string, treeType string) (model.Tree, error) {
	formattedKey := fmt.Sprintf("%s-%s", key, treeType)

	if item, ok := store.Get(formattedKey); ok {
		return item, nil
	}

	if treeType == enum.BinarySearchTree {
		tree := &BST{
			Key: key,
		}

		store.Set(formattedKey, tree)

		return tree, nil
	}

	if treeType == enum.BinaryTree {
		tree := &BinaryTree{
			Key: key,
		}

		store.Set(formattedKey, tree)

		return tree, nil
	}

	return nil, fmt.Errorf("tree type not found: %s", treeType)
}
