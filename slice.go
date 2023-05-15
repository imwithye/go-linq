package linq

type Slice[T any] []T

func Wrap[T any](source []T) Slice[T] {
	return source
}

func Unwrap[T any](source Slice[T]) []T {
	return source
}

// func Any[T any](source []T) bool {
// 	return len(source) > 0
// }

// func All[T any](source []T, predicate func(T) bool) bool {
// 	for _, v := range source {
// 		if !predicate(v) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func First[T any](source []T) T {
// 	return source[0]
// }

// func FirstOrDefault[T any](source []T, d T) T {
// 	if len(source) > 0 {
// 		return source[0]
// 	}
// 	return d
// }

// func Last[T any](source []T) T {
// 	return source[len(source)-1]
// }

// func LastOrDefault[T any](source []T, d T) T {
// 	if len(source) > 0 {
// 		return source[len(source)-1]
// 	}
// 	return d
// }

// func Where[T any](source []T, predicate func(T) bool) []T {
// 	var result []T
// 	for _, v := range source {
// 		if predicate(v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func SelectMany[T, R any](source []T, selector func(T) []R) []R {
// 	var result []R
// 	for _, v := range source {
// 		result = append(result, selector(v)...)
// 	}
// 	return result
// }
