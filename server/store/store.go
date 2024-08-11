package store

import "mmartinjoo/trees/trees"

var Store = make(map[string]*StoreItem)

type StoreItem struct {
	Value *trees.BinaryTree
	Type string
}