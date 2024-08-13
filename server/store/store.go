package store

import (
	"fmt"
	"mmartinjoo/trees/trees"
)

var store = make(map[string]trees.Tree)

var CurrentUnsavedWriteCommands int = 0
var MaxUnsavedWriteCommands int = 3

func Get(key string) (trees.Tree, bool) {
	item, ok := store[key]

	if ok {
		return item, ok
	} else {
		return nil, ok
	}
}

func Set(key string, tree trees.Tree) {
	store[key] = tree
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