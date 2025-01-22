package pokecache

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	Mux      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	c.CacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
	slog.Info(fmt.Sprintf("Added response to cache: %v\n", key))
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	entry, ok := c.CacheMap[key]
	if !ok {
		slog.Info(fmt.Sprintf("Response not in cache: %v\n", key))
		return nil, false
	}
	slog.Info(fmt.Sprintf("Retrieved response from cache: %v\n", key))
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		t := <-ticker.C
		for key, entry := range c.CacheMap {
			if compare := entry.createdAt.Compare(t.Add(-interval)); compare == -1 {
				c.Mux.Lock()
				delete(c.CacheMap, key)
				c.Mux.Unlock()
			}
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{}
	cache.CacheMap = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return cache
}
