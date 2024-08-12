package factory

import (
	"mmartinjoo/trees/store"
	"mmartinjoo/trees/trees"
)

func Create(key string, treeType string) *trees.BST {
	if item, ok := store.Store[key]; ok {
		return item.Value
	}

	if treeType == "binary_tree" {
		tree := trees.BST{
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