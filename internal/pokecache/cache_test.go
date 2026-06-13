package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
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
			key: "https://pokeapi.co/api/v2/pokemon/1",
			val: []byte("bulbasaurdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected key '%s' to exist in cache", c.key)

			}
			if string(val) != string(c.val) {
				t.Errorf("Expected value '%s' for key '%s', got '%s'", c.val, c.key, val)
			}
		})
	}
}

func TestCacheReapLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	const waitTime = 50 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("key1", []byte("value1"))

	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("Expected key 'key1' to exist in cache")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("key1")
	if ok {
		t.Errorf("Expected key 'key1' to have been reaped from cache")
	}
}

func TestCacheGetMissing(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	_, ok := cache.Get("nonexistent_key")
	if ok {
		t.Errorf("Expected key 'nonexistent_key' to not exist in cache")
	}
}
