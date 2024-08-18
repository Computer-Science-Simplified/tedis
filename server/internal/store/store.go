package store

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/model"
)

var store = make(map[string]model.Tree)

var MaxCapacity = 5

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
	if len(store) > MaxCapacity {
		fmt.Println("Store capacity exceeded. Evicting LRU keys...")

		evictionTargets := lru.GetLeastRecentlyUsed(len(store) - MaxCapacity)

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
