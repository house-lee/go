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

type lfNode struct {
    value interface{}
    next  *lfNode
}

type lfQueue struct {
    root         *lfNode
    tail         *lfNode
    maxItems     uint32
    length       uint32
    maxConsumers uint32
    notEmpty     chan bool
}

const uint32subOne = ^uint32(0)

func NewLockFreeQueue(maxConsumers uint32, maxItems uint32) IQueue {
    q := lfQueue{
        root:         &lfNode{nil, nil},
        maxItems:     maxItems,
        length:       0,
        maxConsumers: maxConsumers,
    }
    q.tail = q.root
    q.notEmpty = make(chan bool, maxConsumers)
    return &q
}

func (q *lfQueue) Enqueue(v interface{}) error {
    if q.maxItems != 0 {
        if atomic.AddUint32(&q.length, 1) > q.maxItems {
            atomic.AddUint32(&q.length, uint32subOne)
            return errors.New("ERR_QUEUE_FULL")
        }
    } else {
        atomic.AddUint32(&q.length, 1)
    }
    newNode := &lfNode{
        value: v,
        next:  nil,
    }
    var t *lfNode
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

func (q *lfQueue) Dequeue() interface{} {
    var head *lfNode
    for {
        head = q.root
        if head.next == nil {
            <- q.notEmpty
            continue
        }
        if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.root)), unsafe.Pointer(head), unsafe.Pointer(head.next)) {
            break
        }
    }
    atomic.AddUint32(&q.length, uint32subOne)
    return head.next.value

}

func (q *lfQueue) Length() uint32 {
    length := atomic.LoadUint32(&q.length)
    if length > q.maxItems {
        length = q.maxItems
    }
    return length
}
