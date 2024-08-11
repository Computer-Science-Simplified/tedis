package factory

import (
	"mmartinjoo/trees/trees"

	"github.com/gookit/event"
)

var Store = make(map[string]*StoreItem)

type StoreItem struct {
	Value *trees.BinaryTree
	Type string
}

func Create(key string, treeType string) *trees.BinaryTree {
	if item, ok := Store[key]; ok {
		return item.Value
	}

	if treeType == "binary_tree" {
		tree := trees.BinaryTree{
			Key: key,
		}

		Store[key] = &StoreItem{
			Value: &tree,
			Type: "binary_tree",
		}

		event.MustFire("key_created", event.M{
			"key": key,
		})

		return &tree
	}

	panic("Type not found")
}