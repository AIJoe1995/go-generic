package slice

import "go-generic/set"

func SliceToSet[T comparable](src []T) *set.MapSet[T] {
	length := len(src)
	set := set.NewMapSet[T](length)
	for _, elem := range src {
		set.Add(elem)
	}
	return set
}
