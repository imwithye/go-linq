package linq

func Equal[T comparable](lfs []T, rhs []T) bool {
	if len(lfs) != len(rhs) {
		return false
	}
	for i, v := range lfs {
		if v != rhs[i] {
			return false
		}
	}
	return true
}
