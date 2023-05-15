package linq

func Where[T any](source []T, predicate func(int, T) bool) []T {
	var result []T
	for idx, v := range source {
		if predicate(idx, v) {
			result = append(result, v)
		}
	}
	return result
}

func (s Slice[T]) Where(predicate func(int, T) bool) Slice[T] {
	return Where(s, predicate)
}

func All[T any](source []T, predicate func(int, T) bool) bool {
	for idx, v := range source {
		if !predicate(idx, v) {
			return false
		}
	}
	return true
}

func (s Slice[T]) All(predicate func(int, T) bool) bool {
	return All(s, predicate)
}
