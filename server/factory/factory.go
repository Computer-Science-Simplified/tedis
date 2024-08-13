package factory

import (
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) trees.Tree {
	if item, ok := store.Get(key); ok {
		return item
	}

	if treeType == trees.BinarySearchTree {
		tree := &trees.BST{
			Key: key,
		}

		store.Set(key, tree)

		return tree
	}

	panic("Type not found")
}