package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	storedVal map[string]cacheEntry
	mu        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{}
	//cache.reapLoop(interval)
	return cache
}
