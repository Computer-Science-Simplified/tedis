package store

import (
	"fmt"
	"mmartinjoo/tedis/internal/model"
)

var store = make(map[string]model.Tree)

var CurrentUnsavedWriteCommands int = 0
var MaxUnsavedWriteCommands int = 3

var maxCapacity = 5

func Get(key string) (model.Tree, bool) {
	item, ok := store[key]

	return item, ok
}

func Set(key string, tree model.Tree) {
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
