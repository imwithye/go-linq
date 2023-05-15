package linq

func SelectMany[T, R any](source []T, selector func(T) []R) []R {
	var result []R
	for _, v := range source {
		result = append(result, selector(v)...)
	}
	return result
}
