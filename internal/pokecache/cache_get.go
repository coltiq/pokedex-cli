package pokecache

func (c *Cache) Get(pageURL string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
	entry, ok := c.cache[pageURL]
    return entry.val, ok
}
