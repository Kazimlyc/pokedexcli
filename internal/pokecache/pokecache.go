package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
	cache map[string]cacheEntry
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Lock()
}

func (c *Cache) Get(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Lock()
}

// TODO
func (c *Cache) reapLoop(interval time.Duration) {

}
