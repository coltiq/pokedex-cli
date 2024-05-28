package pokecache

import "time"

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		checkTime := time.Now()
		ticker := time.NewTicker(interval)
		del := make(chan bool)

		go func() {
			for {
				select {
				case <-del:
					c.DeleteEntries(checkTime, interval)
					return
				case <-ticker.C:
					continue
				}
			}
		}()

		time.Sleep(interval)
		ticker.Stop()
		del <- true
	}
}

func (c *Cache) DeleteEntries(checkTime time.Time, interval time.Duration) {
	for pageURL, entry := range c.cache {
		if entry.createdAt.Sub(checkTime) >= interval {
			delete(c.cache, pageURL)
		}
	}
}
