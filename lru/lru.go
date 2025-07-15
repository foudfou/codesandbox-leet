package lru

import "fmt"

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
	h map[int]int
	// l   Node[int]
	l Node
}

func NewLRUCache(cap int) LRUCache {
	l := Node{}
	l.next = &l
	l.prev = &l

	return LRUCache{
		cap: cap,
		// FIXME should the map values be references to Nodes?
		h: make(map[int]int, cap),
		l: l,
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.h[key]
	if !ok {
		return -1
	}

	// Update hash: put key to beginning

	return v
}

// Update
func (c *LRUCache) Put(key, value int) {
	n := c.addNode(key, value)

	c.h[key] = value // FIXME = n

	fmt.Printf("len=%d, cap=%d\n", len(c.h), c.cap)
	if len(c.h) >= c.cap { // evict from end
		fmt.Println("EVICT!")
		last := c.l.prev
		delete(c.h, last.key)

		c.removeNode(last)
	}
}

// TODO make generic
//
//	type Node[K comparable] struct {
//		prev *Node[K]
//		next *Node[K]
//		key  K
//	}
type Node struct {
	prev *Node
	next *Node
	key  int
}

// Inserts at beginning
func (c *LRUCache) addNode(key, value int) *Node {
	n := Node{key: key}

	n.next = c.l.next
	n.prev = &c.l
	c.l.next.prev = &n
	c.l.next = &n

	return &n
}

func (c *LRUCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
