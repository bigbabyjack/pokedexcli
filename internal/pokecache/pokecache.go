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
	mu        *sync.RWMutex
	createdAt time.Time
	cache     map[string]cacheEntry
	duration  time.Duration
}

// NewCache creates a new Cache
// Returns a pointer to the new Cache
func NewCache(d time.Duration) *Cache {
	var sync = &sync.RWMutex{}
	c := &Cache{
		mu:        sync,
		createdAt: getCreatedAt(),
		cache:     make(map[string]cacheEntry),
		duration:  d,
	}
	fmt.Println("cache created")
	c.readLoop()

	return c
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

func (c *Cache) cleanup(time time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("cleaning up cache")
	for key, entry := range c.cache {
		if time.Sub(entry.createdAt) > c.duration {
			c.Delete(key)
		}
	}
	return nil

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cache[key]; !ok {
		return fmt.Errorf("Failed deleting key: %s", key)
	}
	delete(c.cache, key)
	return nil

}

func (c *Cache) readLoop() {
	fmt.Println("readloop started")
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		c.cleanup(getCreatedAt())
	}
}
