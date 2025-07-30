package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Add(string, []byte)
	Get(string) ([]byte, bool)
}

type CacheManager struct {
	Entries map[string]cacheEntry
	mux     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := &CacheManager{
		Entries: make(map[string]cacheEntry),
		mux:     sync.Mutex{},
	}

	ticker := time.NewTicker(interval)
	go cache.reapLoop(ticker)

	return cache
}

func (c *CacheManager) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.Entries[key] = entry
}

func (c *CacheManager) Get(key string) ([]byte, bool) {
	if entry, ok := c.Entries[key]; ok {
		return entry.val, true
	}

	return nil, false
}

func (c *CacheManager) reapLoop(ticker *time.Ticker) {
	for t := range ticker.C {
		c.mux.Lock()

		for key, entry := range c.Entries {
			if entry.createdAt.Before(t) {
				delete(c.Entries, key)
			}
		}

		c.mux.Unlock()
	}
}
