package list

import (
	"go-generic/errs"
	"go-generic/slice"
)

// 对slice的简单封装

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

func NewArrayListFromSlice[T any](slice []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: slice,
	}
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	length := a.Len()
	if index < 0 || index >= length {
		return zero, errs.NewErrIndexOutOfRange(length, index)
	}
	return a.vals[index], nil
}

// Get Set Append Add Delete (shrink) Len Cap

func (a *ArrayList[T]) Set(index int, val T) error {
	length := a.Len()
	if index < 0 || index >= length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = val
	return nil
}

func (a *ArrayList[T]) Append(val ...T) error {
	a.vals = append(a.vals, val...)
	return nil
}

func (a *ArrayList[T]) AddAt(index int, val T) error {

	vals, err := slice.AddAtIndex[T](a.vals, val, index)
	if err != nil {
		return err
	}
	a.vals = vals
	return nil
}

func (a *ArrayList[T]) DeleteAt(index int) (T, error) {

	vals, val, err := slice.DeleteAtIndex(a.vals, index)
	if err != nil {
		return val, err
	}
	a.vals = vals
	return val, nil
}
