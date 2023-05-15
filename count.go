package linq

func Count[T any](source []T) int {
	return len(source)
}

func (s Slice[T]) Count() int {
	return Count(s)
}
