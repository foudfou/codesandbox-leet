package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("lru", func(t *testing.T) {
		lru := NewLRUCache(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		AssertEqual(t, lru.String(), "[2:2,1:1]")
		AssertEqual(t, lru.Get(1), 1) // returns 1
		AssertEqual(t, lru.String(), "[1:1,2:2]")
		lru.Put(3, 3) // evicts key 2
		AssertEqual(t, lru.String(), "[3:3,1:1]")
		AssertEqual(t, lru.Get(2), -1) // returns -1 (not found)
		lru.Put(4, 4)                  // evicts key 1
		AssertEqual(t, lru.String(), "[4:4,3:3]")
		AssertEqual(t, lru.Get(1), -1) // returns -1 (not found)
		AssertEqual(t, lru.Get(3), 3)  // returns 3
		AssertEqual(t, lru.String(), "[3:3,4:4]")
		AssertEqual(t, lru.Get(4), 4) // returns 4
		AssertEqual(t, lru.String(), "[4:4,3:3]")
	})

	t.Run("lru Put updates", func(t *testing.T) {
		lru := NewLRUCache(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(1, 3)
		AssertEqual(t, lru.String(), "[1:3,2:2]")
	})

	t.Run("lru wrong cap", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()
		NewLRUCache(0)
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
