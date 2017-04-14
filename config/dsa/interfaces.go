package dsa

type Queue interface {
    Enqueue(v interface{}) error
    Dequeue() (interface{}, error)
}
