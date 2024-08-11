package listeners

import "mmartinjoo/trees/store"

func EvictOldKeys(data map[string]any, lru store.LRU) {
	key, _ := data["key"].(string)

	_, err := lru.Get(key)

	if err != nil {
		lru.Put(key)
	}

	store.Evict(lru)
}