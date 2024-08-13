package factory

import (
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) *trees.BST {
	if item, ok := store.Get(key); ok {
		return item.Value
	}

	if treeType == trees.BinarySearchTree {
		tree := trees.BST{
			Key: key,
		}

		store.Set(key, &tree, trees.BinarySearchTree)

		return &tree
	}

	panic("Type not found")
}