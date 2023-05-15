package linq

func Select[T, R any](source []T, selector func(int, T) R) []R {
	var result []R
	for i, v := range source {
		result = append(result, selector(i, v))
	}
	return result
}
