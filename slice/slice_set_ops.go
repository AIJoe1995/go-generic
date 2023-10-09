package slice

import "go-generic/set"

func DiffSlice[T comparable](src, dst []T) []T {
	srcSet := SliceToSet(src)
	dstSet := SliceToSet(dst)
	setdiff := set.DiffSet[T](srcSet, dstSet)
	res := setdiff.Keys()
	return res
}

func UnionSlice[T any](slices ...[]T) []T {
	length := 0
	for _, slice := range slices {
		length += len(slice)
	}
	res := make([]T, length)
	for _, slice := range slices {
		for _, elem := range slice {
			res = append(res, elem)
		}
	}
	return res
}
