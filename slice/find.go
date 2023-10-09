package slice

func Exist[T comparable](src []T, element T) bool {
	for _, elem := range src {
		if elem == element {
			return true
		}
	}
	return false
}
