package slice

import "go-generic/errs"

func AddAtIndex[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	src = append(src, element)
	if index == length {
		return src, nil
	}
	for i := len(src) - 1; i > index; i-- {
		src[i] = src[i-1]
	}
	src[index] = element
	return src, nil

}
