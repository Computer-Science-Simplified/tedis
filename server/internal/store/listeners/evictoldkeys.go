package listeners

import (
	store2 "mmartinjoo/trees/internal/store"
)

func EvictOldKeys(data map[string]any, lru *store2.LRU) {
	key, _ := data["key"].(string)

	_, err := lru.Get(key)

	if err != nil {
		lru.Put(key)
	}

	store2.Evict(lru)
}
