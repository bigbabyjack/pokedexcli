package pokecache

import (
	"errors"
	"fmt"
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
	var sync = sync.Mutex{}
	c := Cache{
		mu:        &sync,
		createdAt: getCreatedAt(),
		cache:     make(map[string]cacheEntry),
		duration:  d,
	}
	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("adding key", key)

	c.cache[key] = cacheEntry{
		createdAt: getCreatedAt(),
		val:       val,
	}
	return errors.New("Unable to add key %s")

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reap() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cache {
		if getCreatedAt().Sub(entry.createdAt) > c.duration {
			delete(c.cache, key)
		}
	}
	return nil
}

func (c *Cache) reapLoop() {
	fmt.Println("readloop started")
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		c.reap()
	}
}
