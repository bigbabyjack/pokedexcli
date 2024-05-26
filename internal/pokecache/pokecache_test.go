package pokecache

import (
	"fmt"
	"testing"
	"time"
)

// Test the NewCache function
func TestNewCache(t *testing.T) {
	cache := NewCache(time.Minute * 7)

	// Check if cache is not nil
	if cache == nil {
		t.Fatal("NewCache() returned nil")
	}

	// Check if createdAt is set correctly (within a reasonable range)
	now := time.Now()
	if cache.createdAt.Before(now.Add(-time.Second)) || cache.createdAt.After(now) {
		t.Fatalf("Expected createdAt to be around %v, got %v", now, cache.createdAt)

	}

	// Check if cache slice is initialized and empty
	if len(cache.cache) != 0 {
		t.Fatalf("Expected empty cache slice, got %d entries", len(cache.cache))
	}
}
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(time.Nanosecond * 100)

	// Add an entry to the cache
	cache.Add("key1", []byte("value1"))

	// Delete the entry from the cache
	cache.Delete("key1")

	// Check if the entry is deleted
	if _, ok := cache.cache["key1"]; ok {
		t.Fatalf("Expected cache entry to be deleted")
	}
}

func TestCleanup(t *testing.T) {
	cache := NewCache(time.Millisecond * 50)

	// Add an entry to the cache
	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))

	// Call the cleanup function
	time.Sleep(time.Millisecond * 100)
	cache.cleanup(time.Now())

	// Check if the entry is deleted
	if _, ok := cache.cache["key1"]; ok {
		t.Fatalf("Expected cache entry to be deleted")
	}
	if _, ok := cache.cache["key2"]; ok {
		t.Fatalf("Expected cache entry to be deleted")
	}

	// Add an entry to the cache
	cache.Add("key3", []byte("value3"))
	cache.Add("key4", []byte("value4"))

	// Call the cleanup function
	cache.cleanup(time.Now())

	// Check if the entry is deleted
	if _, ok := cache.cache["key3"]; !ok {
		t.Fatalf("Expected cache entry to be cached")
	}
	if _, ok := cache.cache["key4"]; !ok {
		t.Fatalf("Expected cache entry to be cached")
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
