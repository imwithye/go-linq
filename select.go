package linq

func Select[T, R any](source []T, selector func(T) R) []R {
	result := make([]R, len(source))
	for i, v := range source {
		result[i] = selector(v)
	}
	return result
}
