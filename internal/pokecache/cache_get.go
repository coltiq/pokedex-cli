package pokecache

func (c *Cache) Get(pageURL string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cache[pageURL]
	c.mu.Unlock()
    return entry.val, ok
}
