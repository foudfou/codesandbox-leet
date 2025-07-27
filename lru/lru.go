package lru

import (
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
type LRUCache struct {
	cap int
	// len = len(map)
	l Node
	h map[int]*Node
}

func NewLRUCache(cap int) *LRUCache {
	if cap <= 0 {
		panic("capacity must be positive")
	}

	c := &LRUCache{
		cap,
		Node{},
		make(map[int]*Node),
	}

	c.l.prev = &c.l
	c.l.next = &c.l

	return c
}

func (c *LRUCache) Put(k, v int) {
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
			delete(c.h, last.val)
		}

		n = &Node{key: k, val: v}
		insertNode(&c.l, n)

		c.h[k] = n
	}
}

func (c *LRUCache) Get(k int) int {
	n, ok := c.h[k]
	if !ok {
		return -1
	}

	// move to front
	deleteNode(n)
	insertNode(&c.l, n)

	return n.val
}

// Useful for debugging or testing
func (c *LRUCache) String() string {
	if c.l.next == &c.l {
		return "[]"
	}

	nodes := []string{}
	for p := c.l.next; p != &c.l; p = p.next {
		nodes = append(nodes, fmt.Sprintf("%v:%v", p.key, p.val))
	}
	return "[" + strings.Join(nodes, ",") + "]"
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
	key  int
	val  int
}

// Insert node n to front of list l
func insertNode(l, n *Node) {
	n.next = l.next
	n.prev = l
	l.next.prev = n
	l.next = n
}

func deleteNode(n *Node) {
	n.next.prev = n.prev
	n.prev.next = n.next
}
