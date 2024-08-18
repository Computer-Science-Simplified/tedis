package listeners

import "github.com/Computer-Science-Simplified/tedis/server/internal/store"

func PutLRUItem(lru *store.LRU, key string) error {
	if lru.Exists(key) {
		_, err := lru.Get(key)
		if err != nil {
			return err
		}
	} else {
		lru.Put(key)
	}

	return nil
}
