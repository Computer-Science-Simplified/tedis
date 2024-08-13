package factory

import (
	"fmt"
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) (trees.Tree, error) {
	if item, ok := store.Get(key); ok {
		return item, nil
	}

	if treeType == trees.BinarySearchTree {
		tree := &trees.BST{
			Key: key,
		}

		store.Set(key, tree)

		return tree, nil
	}

	return nil, fmt.Errorf("tree type not found: %s", treeType)
}