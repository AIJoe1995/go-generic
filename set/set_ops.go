package set

func DiffSet[T comparable](s1, s2 *MapSet[T]) *MapSet[T] {
	// return elements in s1 but not in s2
	res := NewMapSet[T](len(s1.m))
	for elem := range s1.m {
		if _, ok := s2.m[elem]; !ok {
			res.Add(elem)
		}
	}
	return res

}

func UnionSet[T comparable](sets ...*MapSet[T]) *MapSet[T] {
	length := 0
	for _, set := range sets {
		length += len(set.m)
	}
	res := NewMapSet[T](length) // how to determine the size of MapSet
	for _, set := range sets {
		for elem := range set.m {
			res.Add(elem)
		}
	}
	//todo resize the MapSet
	return res
}

func IntersectionSet[T comparable](s1, s2 *MapSet[T]) *MapSet[T] {
	l1 := len(s1.m)
	l2 := len(s2.m)
	length := 0
	if l1 > l2 {
		length = l1
	} else {
		length = l2
	}
	res := NewMapSet[T](length)

	for elem := range s1.m {
		if _, ok := s1.m[elem]; ok {
			res.Add(elem)
		}
	}
	return res
}
