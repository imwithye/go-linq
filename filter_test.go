package linq

import (
	"testing"
)

func TestWhere(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := []int{2, 4}
	got := Where(source, func(i int, v int) bool {
		return v%2 == 0
	})

	if !Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
