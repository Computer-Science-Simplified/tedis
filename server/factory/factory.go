package factory

import (
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) *trees.BinaryTree {
	if item, ok := store.Store[key]; ok {
		return item.Value
	}

	if treeType == "binary_tree" {
		tree := trees.BinaryTree{
			Key: key,
		}

		store.Store[key] = &store.StoreItem{
			Value: &tree,
			Type: "binary_tree",
		}

		return &tree
	}

	panic("Type not found")
}