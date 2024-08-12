package factory

import (
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) *trees.BST {
	if item, ok := store.Store[key]; ok {
		return item.Value
	}

	if treeType == trees.BinarySearchTree {
		tree := trees.BST{
			Key: key,
		}

		store.Store[key] = &store.StoreItem{
			Value: &tree,
			Type: trees.BinarySearchTree,
		}

		return &tree
	}

	panic("Type not found")
}