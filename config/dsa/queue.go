package dsa

/**
 * Utility: Lock Free Queue
 * Package DSA (Data Structures and Algorithms)
 */

type node struct {
	value interface{}
	next  *node
}

type queue struct {
	head     *node
	tail     *node
	maxItems uint32
}
