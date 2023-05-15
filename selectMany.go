package linq

func SelectMany[T, R any](source []T, selector func(int, T) []R) []R {
	var result []R
	for idx, v := range source {
		result = append(result, selector(idx, v)...)
	}
	return result
}
