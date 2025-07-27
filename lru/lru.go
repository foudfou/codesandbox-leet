package lru

import (
	"errors"
	"fmt"
	"strings"
)

// LRU Cache
//
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// Implement the LRUCache class:
//
//     LRUCache(capacity int) - Initialize with positive size capacity
//     Get(key int) int - Return the value of the key if it exists, otherwise return -1. Move the key to most recent.
//     Put(key, value int) - Update the value if key exists, otherwise add the key-value pair. If adding causes capacity to be exceeded, remove the least recently used key. Move the key to most recent.
//
// Both Get and Put must run in O(1) average time complexity.

// LRU cache is a linked list to hold ordered values + a map pointing to the list items for O(1) retrieval
type LRUCache[K comparable, V any] struct {
	cap int
	// len = len(map)
	h map[K]*Node[K, V]
	l Node[K, V]
}

func NewLRUCache[K comparable, V any](cap int) *LRUCache[K, V] {
	if cap <= 0 {
		panic("capacity must be positive")
	}

	c := &LRUCache[K, V]{
		cap,
		make(map[K]*Node[K, V]),
		Node[K, V]{},
	}

	c.l.prev = &c.l
	c.l.next = &c.l

	return c
}

func (c *LRUCache[K, V]) Put(k K, v V) {
	n, ok := c.h[k]
	if ok { // update
		n.val = v

		// move node to front
		deleteNode(n)
		insertNode(&c.l, n)
	} else { // insert
		if len(c.h) >= c.cap {
			last := c.l.prev
			deleteNode(last)
			delete(c.h, last.key)
		}

		n = &Node[K, V]{key: k, val: v}
		insertNode(&c.l, n)

		c.h[k] = n
	}
}

var ErrNotFound = errors.New("Resource was not found")

// Returns value for k, or the zero value with ErrNotFound if not present.
func (c *LRUCache[K, V]) Get(k K) (V, error) {
	n, ok := c.h[k]
	if !ok {
		var zero V
		return zero, ErrNotFound
	}

	// move to front
	deleteNode(n)
	insertNode(&c.l, n)

	return n.val, nil
}

// Useful for debugging or testing
func (c *LRUCache[K, V]) String() string {
	if c.l.next == &c.l {
		return "[]"
	}

	nodes := []string{}
	for p := c.l.next; p != &c.l; p = p.next {
		nodes = append(nodes, fmt.Sprintf("%v:%v", p.key, p.val))
	}
	return "[" + strings.Join(nodes, ",") + "]"
}

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	key  K
	val  V
}

// Insert node n to front of list l
func insertNode[K comparable, V any](l, n *Node[K, V]) {
	n.next = l.next
	n.prev = l
	l.next.prev = n
	l.next = n
}

func deleteNode[K comparable, V any](n *Node[K, V]) {
	n.next.prev = n.prev
	n.prev.next = n.next
}
