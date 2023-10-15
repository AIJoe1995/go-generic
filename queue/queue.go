package queue

import "go-generic/errs"

type node[T any] struct {
	value T
	next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

type SimpleQueue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewSimpleQueue[T any]() *SimpleQueue[T] {
	return &SimpleQueue[T]{}
}

func SliceToSimpleQueue[T any](slice []T) *SimpleQueue[T] {
	sq := NewSimpleQueue[T]()
	for _, n := range slice {
		sq.Enqueue(n)
	}
	return sq
}

func QueueToSlice[T any](que *SimpleQueue[T]) []T {
	length := que.size
	slice := make([]T, 0, length)
	newque := NewSimpleQueue[T]()
	for i := 0; i < length; i++ {
		val, _ := que.Dequeue()
		slice = append(slice, val)
		newque.Enqueue(val)
	}
	que = newque
	return slice
}

func (q *SimpleQueue[T]) Enqueue(t T) error {
	newnode := newNode[T](t)
	if q.Len() == 0 {
		q.head = newnode
		q.tail = newnode
		q.size = 1
		return nil
	}
	q.tail.next = newnode
	q.tail = q.tail.next
	q.size += 1
	return nil
}

func (q *SimpleQueue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errs.NewErrEmptyQueue()
	}

	node := q.head
	return node.value, nil

}

func (q *SimpleQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errs.NewErrEmptyQueue()
	}
	if q.Len() == 1 {
		node := q.head
		// todo 删掉node 释放内存
		q.head.next = nil
		q.tail.next = nil
		q.size = 0
		return node.value, nil
	}
	node := q.head
	q.head = node.next
	q.size -= 1
	return node.value, nil

}

func (q *SimpleQueue[T]) Len() int {
	return q.size
}

func (q *SimpleQueue[T]) IsEmpty() bool {
	return q.size == 0
}
