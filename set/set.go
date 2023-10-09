package set

import "go-generic/errs"

type Set[T comparable] interface {
	Add(key T)
	Remove(key T) error
	Exist(key T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s MapSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}

func (s MapSet[T]) Remove(key T) error {
	_, ok := s.m[key]
	if !ok {
		return errs.NewErrKeyNotExist()
	}
	delete(s.m, key)
	return nil
}

func (s MapSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s MapSet[T]) Keys() []T {
	res := make([]T, 0, len(s.m))
	for key := range s.m {
		res = append(res, key)
	}
	return res
}
