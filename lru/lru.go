package lru

import (
	"fmt"
	"strings"
)

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
//
// LRUCache(capacity int) - Initialize with positive size capacity
// Get(key int) int - Return the value of the key if it exists, otherwise return -1. Move the key to most recent.
// Put(key, value int) - Update the value if key exists, otherwise add the key-value pair. If adding causes capacity to be exceeded, remove the least recently used key. Move the key to most recent.
//
// Both Get and Put must run in O(1) average time complexity.

type LRUCache struct {
	cap int
	// len via len(c.h)
	h map[int]*Node
	l Node
}

func NewLRUCache(cap int) *LRUCache {
	if cap <= 0 {
		panic("capacity must be positive")
	}

	cache := LRUCache{
		cap: cap,
		h:   make(map[int]*Node, cap),
		l:   Node{},
	}

	cache.l.next = &cache.l
	cache.l.prev = &cache.l

	// return by reference to avoid copying and breaking circular reference.
	// FIXME allocating on the heap is best practice anyways
	return &cache
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.h[key]
	if !ok {
		return -1
	}

	c.removeNode(n)
	c.prependNode(n)

	return n.val
}

// Update
func (c *LRUCache) Put(key, val int) {
	if n, ok := c.h[key]; ok {
		n.val = val
		c.removeNode(n)
		c.prependNode(n)
		return
	}

	n := &Node{key: key, val: val}
	c.prependNode(n)
	c.h[key] = n

	// fmt.Printf("len=%d, cap=%d\n", len(c.h), c.cap)
	if len(c.h) > c.cap { // evict from end
		// fmt.Println("EVICT!")
		last := c.l.prev
		delete(c.h, last.key)

		c.removeNode(last)
	}
}

// TODO make generic
//
//	type Node[K comparable, V any] struct {
//	    prev *Node[K, V]
//	    next *Node[K, V]
//	    key  K
//	    val  V
//	}
type Node struct {
	prev *Node
	next *Node
	// key is needed to delete evicted node. Also nice for String().
	key int
	val int
}

// Inserts at beginning
func (c *LRUCache) prependNode(n *Node) {
	n.next = c.l.next
	n.prev = &c.l
	c.l.next.prev = n
	c.l.next = n
}

func (c *LRUCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) String() string {
	// fmt.Printf("l=%p, next=%p, prev=%p\n", &c.l, c.l.next, c.l.prev)
	var elts []string
	for n := c.l.next; n != &c.l; n = n.next {
		elts = append(elts, fmt.Sprintf("%v:%v", n.key, n.val))
	}
	return "[" + strings.Join(elts, ",") + "]"
}
