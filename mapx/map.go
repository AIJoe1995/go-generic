package mapx

func Keys[K comparable, V any](m map[K]V) []K {
	length := len(m)
	res := make([]K, length)
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func Values[K comparable, V any](m map[K]V) []V {
	length := len(m)
	res := make([]V, length)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func UpdateMap[K comparable, V any](m1, m2 map[K]V) map[K]V {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
