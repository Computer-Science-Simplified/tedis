package factory

import "mmartinjoo/trees/trees"

var store = make(map[string]*trees.BinaryTree)

func Create(key string, treeType string) *trees.BinaryTree {
	if tree, ok := store[key]; ok {
		return tree
	}

	if treeType == "binary_tree" {
		tree := trees.BinaryTree{
			Key: key,
		}

		store[key] = &tree

		return &tree
	}

	panic("Type not found")
}