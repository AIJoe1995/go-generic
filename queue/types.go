package queue

type Queue[T any] interface {
	Enqueue(t T) error
	Dequeue() (T, error)
}
