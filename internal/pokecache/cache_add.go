package pokecache

import "time"

func (c *Cache) Add(pageURL string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.cache[pageURL] = entry
	c.mu.Unlock()
}
