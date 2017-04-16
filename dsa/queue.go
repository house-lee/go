package dsa

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

/**
 * Utility: Lock Free Queue
 * Package DSA (Data Structures and Algorithms)
 */

type node struct {
	value interface{}
	next  *node
}

type queue struct {
	root         *node
	tail         *node
	maxItems     uint32
	length       uint32
	maxConsumers uint32
	notEmpty		 chan bool
}

const uint32subOne = ^uint32(0)

func NewQueue(maxConsumers uint32, maxItems uint32) IQueue {
	q := queue{
		root:         &node{nil, nil},
		maxItems:     maxItems,
		length:       0,
		maxConsumers: maxConsumers,
	}
	q.tail = q.root
	q.notEmpty = make(chan bool, maxConsumers)
	return &q
}

func (q *queue) Enqueue(v interface{}) error {
	if q.maxItems != 0 {
		if atomic.AddUint32(&q.length, 1) > q.maxItems {
			atomic.AddUint32(&q.length, uint32subOne)
			return errors.New("ERR_QUEUE_FULL")
		}
	}
	newNode := &node{
		value: v,
		next:  nil,
	}
	var t *node
	for {
		t = q.tail
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&t.next)), nil, unsafe.Pointer(newNode)) {
			break
		}
	}
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(t), unsafe.Pointer(newNode))
	select {
	case q.notEmpty <- true:
	default:

	}
	return nil
}

func (q *queue) Dequeue() interface{} {
	var head *node
	for {
		head = q.root.next
		if head == nil {
			<- q.notEmpty
			continue
		}
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.root.next)), unsafe.Pointer(head), unsafe.Pointer(head.next)) {
			break
		}
	}
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(head), unsafe.Pointer(q.root))
	atomic.AddUint32(&q.length, uint32subOne)
	return head.value
}
