package store

import (
	"errors"
)

type LRU struct {
	Map      map[string]string
	Items    []string
	Capacity int
}

func NewLRU() *LRU {
	return &LRU{
		Capacity: MaxCapacity,
		Map:      make(map[string]string, MaxCapacity),
		Items:    make([]string, MaxCapacity),
	}
}

func (lru *LRU) Put(key string) {
	if len(lru.Items) >= lru.Capacity {
		removedKey := lru.Items[len(lru.Items)-1]

		lru.Items = lru.Items[:len(lru.Items)-1]

		delete(lru.Map, removedKey)
	}

	lru.Items = append([]string{key}, lru.Items...)

	lru.Map[key] = key
}

func (lru *LRU) Get(key string) (string, error) {
	res, ok := lru.Map[key]

	if !ok {
		return "", errors.New("key does not exist")
	}

	lru.update(key)

	return res, nil
}

func (lru *LRU) Count() int {
	return len(lru.Items)
}

func (lru *LRU) GetLeastRecentlyUsed(n int) []string {
	targets := make([]string, n)

	copy(targets, lru.Items[len(lru.Items)-n:])

	return targets
}

func (lru *LRU) Remove(key string) {
	delete(lru.Map, key)

	for idx, value := range lru.Items {
		if value == key {
			lru.Items = append(lru.Items[:idx], lru.Items[idx+1:]...)
			break
		}
	}
}

func (lru *LRU) Exists(key string) bool {
	_, ok := lru.Map[key]

	return ok
}

func (lru *LRU) update(key string) {
	idx := -1

	for i, value := range lru.Items {
		if value == key {
			idx = i
			break
		}
	}

	if idx == -1 {
		return
	}

	lru.Items = append(lru.Items[:idx], lru.Items[idx+1:]...)

	lru.Items = append([]string{key}, lru.Items...)
}
