package linq

func Count[T any](source []T) int {
	return len(source)
}

func (s Slice[T]) Count() int {
	return Count(s)
}

func Any[T any](source []T) bool {
	return len(source) > 0
}

func (s Slice[T]) Any() bool {
	return Any(s)
}
