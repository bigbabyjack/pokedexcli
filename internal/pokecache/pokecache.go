package pokecache

import (
	"errors"
	"fmt"
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
	createdAt time.Time
	cache     map[string]cacheEntry
	duration  time.Duration
}

// NewCache creates a new Cache
// Returns a pointer to the new Cache
func NewCache(d time.Duration) *Cache {
	c := &Cache{
		createdAt: getCreatedAt(),
		cache:     make(map[string]cacheEntry),
		duration:  d,
	}
	// c.realLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) error {
	c.cache[key] = cacheEntry{
		createdAt: getCreatedAt(),
		val:       val,
	}
	return errors.New("Unable to add key %s")

}

func (c *Cache) cleanup(time time.Time) error {
	for key, entry := range c.cache {
		if time.Sub(entry.createdAt) > c.duration {
			c.Delete(key)
		}
	}
	return nil

}

func (c *Cache) Get(key string) ([]byte, error) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, errors.New("Key not found")
	}
	return entry.val, nil
}

func (c *Cache) Delete(key string) error {
	if _, ok := c.cache[key]; !ok {
		return fmt.Errorf("Failed deleting key: %s", key)
	}
	delete(c.cache, key)
	return nil

}

// func (c *Cache) realLoop()
