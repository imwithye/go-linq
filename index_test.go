package linq

import (
	"testing"
)

func TestFirst(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := 1
	got := First(source)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFirstOrDefault(t *testing.T) {
	source := []int{}
	want := 1
	got := FirstOrDefault(source, 1)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLast(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := 5
	got := Last(source)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLastOrDefault(t *testing.T) {
	source := []int{}
	want := 5
	got := LastOrDefault(source, 5)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAt(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	want := 3
	got := At(source, 2)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
