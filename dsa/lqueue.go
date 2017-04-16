package dsa

import (
    "sync"
    "errors"
)

type lNode struct {
	value interface{}
	next  *lNode
}

type lQueue struct {
	root         *lNode
	tail         *lNode
	length       uint32
	maxItems     uint32
	maxConsumers uint32
	opLock       sync.RWMutex
    notEmpty     chan bool
}

func NewQueueWithLock(maxConsumers uint32, maxItems uint32) IQueue {
    q := lQueue{
        root: &lNode{nil,nil},
        maxItems:maxItems,
        length:0,
        maxConsumers:maxConsumers,
    }
    q.tail = q.root
    q.notEmpty = make(chan bool, maxConsumers)
    return &q
}

func (q *lQueue)Enqueue(v interface{}) error  {
    q.opLock.Lock()

    if q.maxItems != 0 && q.length == q.maxItems {
        q.opLock.Unlock()
        return errors.New("ERR_QUEUE_FULL")
    }
    q.length++
    newNode := &lNode{
        value:v,
        next:nil,
    }
    q.tail.next = newNode
    q.tail = newNode
    q.opLock.Unlock()
    select {
    case q.notEmpty <- true:
    default:
    }
    return nil
}

func (q *lQueue)Dequeue() interface{} {
    var head *lNode
    for {
        q.opLock.Lock()
        head = q.root
        if head.next == nil {
            q.opLock.Unlock()
            <- q.notEmpty
            continue
        }
        q.root = head.next
        q.length--
        q.opLock.Unlock()
        break
    }
    return head.next.value
}

func (q *lQueue) Length() uint32  {
    q.opLock.RLock()
    defer q.opLock.RUnlock()
    return q.length
}