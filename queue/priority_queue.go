package queue

import (
	"go-generic/errs"
)

type Comparator[T any] func(src T, dst T) int

type PriorityQueue[T any] struct {
	data     []T
	compare  Comparator[T]
	capacity int
}

func NewPriorityQueue[T any](capacity int, compare Comparator[T]) *PriorityQueue[T] {
	slicecap := capacity
	if capacity < 1 {
		capacity = 0
		slicecap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		compare:  compare,
		data:     make([]T, 0, slicecap),
	}
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.data)
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	var zero T
	if pq.Len() == 0 {
		return zero, errs.NewErrEmptyQueue()
	}
	val := pq.data[0]
	return val, nil
}

func (pq *PriorityQueue[T]) IsFull() bool {
	return pq.Len() == pq.capacity
}

func (pq *PriorityQueue[T]) Enqueue(elem T) error {
	if pq.IsFull() {
		return errs.NewErrQueueFull()
	}
	pq.data = append(pq.data, elem)
	pq.siftup(elem, pq.Len()-1)
	return nil
}

func (pq *PriorityQueue[T]) siftup(elem T, idx int) {
	parent_idx := (idx - 1) / 2
	current_idx := idx
	for current_idx > 0 && pq.compare(elem, pq.data[parent_idx]) < 0 {
		pq.data[current_idx] = pq.data[parent_idx]
		current_idx, parent_idx = parent_idx, (parent_idx-1)/2
	}
	pq.data[current_idx] = elem
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	var zero T
	if pq.Len() == 0 {
		return zero, errs.NewErrEmptyQueue()
	}
	res := pq.data[0]
	end := pq.data[pq.Len()-1]
	pq.data[0] = end
	pq.siftdown(end, 0, pq.Len()-1)
	return res, nil
}

func (pq *PriorityQueue[T]) siftdown(val T, start_idx int, end_idx int) {
	parent_idx := start_idx
	child_idx := parent_idx*2 + 1 // left child
	for child_idx < end_idx {
		if child_idx+1 < end_idx && pq.compare(pq.data[child_idx+1], pq.data[child_idx]) < 0 {
			// 存在右节点且右节点小
			child_idx += 1 // 右节点
		}
		if pq.compare(val, pq.data[child_idx]) < 0 {
			break
		}
		pq.data[child_idx] = pq.data[parent_idx]
		parent_idx, child_idx = child_idx, 2*parent_idx+1
	}
	pq.data[child_idx] = val
	pq.data = pq.data[:end_idx]

}
