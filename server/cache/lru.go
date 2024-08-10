package cache

import "errors"

type LRU struct {
	Map map[string]string
	Items []string
	Capacity int
}

func NewLRU(capacity int) LRU {
	return LRU{
		Capacity: capacity,
		Map: make(map[string]string, capacity),
		Items: make([]string, capacity),
	}
}

func (lru *LRU) Put(key string) {
	if len(lru.Items) >= lru.Capacity {
		removedKey := lru.Items[len(lru.Items) - 1]

		lru.Items = lru.Items[:len(lru.Items) - 1]

		delete(lru.Map, removedKey)
	}

	lru.Items = append([]string{key}, lru.Items...)

	lru.Map[key] = key
}

func (lru *LRU) Get(key string) (string, error) {
	_, exists := lru.Map[key]

	if !exists {
		return "", errors.New("key does not exist")
	}

	lru.update(key)

	return lru.Map[key], nil
}

func (lru *LRU) update(key string) {
	idx := -1

	for i, value := range lru.Items {
		if value == key {
			idx = i
			break;
		}
	}

	if idx == -1 {
		return
	}

	lru.Items = append(lru.Items[:idx], lru.Items[idx+1:]...)

	lru.Items = append([]string{key}, lru.Items...)
}
