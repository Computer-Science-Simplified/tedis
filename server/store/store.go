package store

import (
	"fmt"
	"mmartinjoo/trees/trees"
)

var store = make(map[string]*StoreItem)

var CurrentUnsavedWriteCommands int = 0
var MaxUnsavedWriteCommands int = 3

type StoreItem struct {
	Value trees.Tree
	Type string
}

func Get(key string) (*StoreItem, bool) {
	item, ok := store[key]

	if ok {
		return item, ok
	} else {
		return nil, ok
	}
}

func Set(key string, tree *trees.BST, treeType string) {
	store[key] = &StoreItem{
		Value: tree,
		Type: treeType,
	}
}

func Len() int {
	return len(store)
}

func Keys() []string {
	keys := make([]string, 0)

	for key := range store {
		keys = append(keys, key)
	}

	return keys
}

var maxCapacity = 5

func Evict(lru *LRU) {
	if len(store) > maxCapacity {
		fmt.Println("Store capacity exceeded. Evicting LRU keys...")

		evictionTargets := lru.GetLeastRecentlyUsed(len(store) - maxCapacity)

		numberOfEvictedKeys := 0

		fmt.Printf("Eviction targets: %v\n", evictionTargets)

		for _, evictionTarget := range evictionTargets {
			delete(store, evictionTarget)
			lru.Remove(evictionTarget)

			numberOfEvictedKeys++
		}

		fmt.Printf("Evicted %d keys\n", numberOfEvictedKeys)
	}
}

func ShouldPersist() bool {
	return CurrentUnsavedWriteCommands%MaxUnsavedWriteCommands == 0
}