package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func EvictOldKeys(data map[string]any, lru *store.LRU) {
	key, _ := data["key"].(string)

	_, err := lru.Get(key)

	if err != nil {
		lru.Put(key)
	}

	store.Evict(lru)
}
