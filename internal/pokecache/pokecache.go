package pokecache

import (
	"errors"
	"sync"
	"time"
)

func getCreatedAt() time.Time {
	return time.Now()
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu        *sync.Mutex
	createdAt time.Time
	cache     map[string]cacheEntry
	duration  time.Duration
}

// NewCache creates a new Cache
// Returns a pointer to the new Cache
func NewCache(d time.Duration) *Cache {
	c := &Cache{
		mu:        &sync.Mutex{},
		createdAt: time.Now().UTC(),
		cache:     make(map[string]cacheEntry),
		duration:  d,
	}
	go c.reapLoop(c.duration)

	return c
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: getCreatedAt(),
		val:       val,
	}
	return errors.New("Unable to add key %s")

}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reap(interval time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeCutoff := time.Now().UTC().Add(-interval)

	for key, entry := range c.cache {
		if (entry.createdAt).Before(timeCutoff) {
			delete(c.cache, key)
		}
	}
	return nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		c.reap(interval)
	}
}
