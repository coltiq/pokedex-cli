package pokecache

import "time"

func (c *Cache) Add(pageURL string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[pageURL] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
