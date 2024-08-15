package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func EvictOldKeys(mruKey string, lru *store.LRU) {
	// By accessing it, LRU puts the most recently used key at the beginning
	_, err := lru.Get(mruKey)

	// If it's not found for some reason we put it at the end
	if err != nil {
		lru.Put(mruKey)
	}

	store.Evict(lru)
}
