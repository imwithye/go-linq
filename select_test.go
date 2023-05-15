package linq

import (
	"strconv"
	"testing"
)

func TestSelect(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	selector := func(idx int, v int) string {
		return strconv.Itoa(v * 2)
	}
	want := []string{"2", "4", "6", "8", "10"}
	got := Select(source, selector)

	if !Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
