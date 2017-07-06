package dsa

type IQueue interface {
    Enqueue(v interface{}) error
    Dequeue() interface{}
    Length() uint32
}
