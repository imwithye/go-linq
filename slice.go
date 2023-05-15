package linq

type Slice[T any] []T

func Wrap[T any](source []T) Slice[T] {
	return source
}

func Unwrap[T any](source Slice[T]) []T {
	return source
}
