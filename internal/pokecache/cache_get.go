package pokecache

func (c *Cache) Get(pageURL string) ([]byte, bool) {
	c.mu.Lock()
	if entry, ok := c.storedVal[pageURL]; ok {
		return entry.val, true
	}
	c.mu.Unlock()
	return nil, false
}
