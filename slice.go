package linq

type Slice[T any] []T

func Wrap[T any](source []T) Slice[T] {
	return source
}

func Unwrap[T any](source Slice[T]) []T {
	return source
}

// func All[T any](source []T, predicate func(T) bool) bool {
// 	for _, v := range source {
// 		if !predicate(v) {
// 			return false
// 		}
// 	}
// 	return true
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
