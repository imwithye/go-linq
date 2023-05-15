package linq

import (
	"strconv"
	"testing"
)

func TestSelectMany(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	selector := func(idx int, v int) []string {
		return []string{strconv.Itoa(v * 2), strconv.Itoa(v*2 + 1)}
	}
	want := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}
	got := SelectMany(source, selector)

	if !Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
