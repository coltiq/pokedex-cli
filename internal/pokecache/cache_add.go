package pokecache

import "time"

func (c *Cache) Add(pageURL string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.storedVal[pageURL] = entry
	c.mu.Unlock()
}
