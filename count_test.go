package linq

import (
	"testing"
)

func TestCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := 5
	got := Count(source)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAny(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := true
	got := Any(source)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
