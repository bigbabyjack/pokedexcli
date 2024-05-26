package pokecache

import (
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

// Test the Add function
func TestAdd(t *testing.T) {
	cache := NewCache(time.Minute * 7)

	// Add an entry to the cache
	cache.Add("key1", []byte("value1"))

	// Check if cache has one entry
	if len(cache.cache) != 1 {
		t.Fatalf("Expected 1 entry in cache, got %d", len(cache.cache))
	}

	// Check the contents of the entry
	entry := cache.cache["key1"]
	if string(entry.val) != "value1" {
		t.Fatalf("Expected cache entry value 'value1', got '%s'", string(entry.val))
	}

	// Check if createdAt is set correctly (within a reasonable range)
	now := time.Now()
	if entry.createdAt.Before(now.Add(-time.Second)) || entry.createdAt.After(now) {
		t.Fatalf("Expected createdAt to be around %v, got %v", now, entry.createdAt)
	}
}

// Test the Get function
func TestGet(t *testing.T) {
	cache := NewCache(time.Minute * 7)

	// Add an entry to the cache
	cache.Add("key1", []byte("value1"))

	// Get the entry from the cache
	val, err := cache.Get("key1")

	// Check if the entry is returned correctly
	if err != nil {
		t.Fatalf("Error getting entry: %v", err)
	}
	if string(val) != "value1" {
		t.Fatalf("Expected cache entry value 'value1', got '%s'", string(val))
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
