package store

import (
	"fmt"
	"mmartinjoo/trees/trees"
)

var Store = make(map[string]*StoreItem)

type StoreItem struct {
	Value *trees.BinaryTree
	Type string
}

var maxCapacity = 5

func Evict(lru LRU) {
	if len(Store) >= maxCapacity {
		fmt.Println("Store capacity exceeded. Evicting LRU keys...")

		evictionTargets := lru.GetLeastRecentlyUsed(len(Store) - maxCapacity)

		numberOfEvictedKeys := 0

		fmt.Printf("Eviction targets: %v\n", evictionTargets)

		for _, evictionTarget := range evictionTargets {
			delete(Store, evictionTarget)
			lru.Remove(evictionTarget)

			numberOfEvictedKeys++
		}

		fmt.Printf("Evicted %d keys\n", numberOfEvictedKeys)
	}
}