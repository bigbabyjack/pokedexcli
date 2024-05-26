package pokecache

import (
	"fmt"
	"testing"
	"time"
)

// Test the NewCache function
func TestNewCache(t *testing.T) {
	cache := NewCache(time.Minute * 7)

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
