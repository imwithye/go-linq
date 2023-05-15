package linq

func First[T any](source []T) T {
	return source[0]
}

func (s Slice[T]) First() T {
	return First(s)
}

func FirstOrDefault[T any](source []T, d T) T {
	if len(source) > 0 {
		return source[0]
	}
	return d
}

func (s Slice[T]) FirstOrDefault(d T) T {
	return FirstOrDefault(s, d)
}

func Last[T any](source []T) T {
	return source[len(source)-1]
}

func (s Slice[T]) Last() T {
	return Last(s)
}

func LastOrDefault[T any](source []T, d T) T {
	if len(source) > 0 {
		return source[len(source)-1]
	}
	return d
}

func (s Slice[T]) LastOrDefault(d T) T {
	return LastOrDefault(s, d)
}

func At[T any](source []T, index int) T {
	return source[index]
}

func (s Slice[T]) At(index int) T {
	return At(s, index)
}
