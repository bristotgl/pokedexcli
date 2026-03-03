package pokecache

import (
	"slices"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := map[string]struct {
		key string
		val []byte
	}{
		"random data": {
			key: "https://example.com",
			val: []byte("testdata"),
		},
		"empty data": {
			key: "https://example.com/path",
			val: []byte(""),
		},
		"empty key": {
			key: "",
			val: []byte("testdata"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			cachedVal, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key '%s'", c.key)
				return
			}

			if !slices.Equal(cachedVal, c.val) {
				t.Errorf("Expected cached value '%v' to be equal to input '%v'", string(cachedVal), string(c.val))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + time.Millisecond
	const key = "https://example.com"

	cache := NewCache(baseTime)
	cache.Add(key, []byte("test data"))

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to find key '%s'", key)
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("Expected to not find key '%s'", key)
		return
	}

}
