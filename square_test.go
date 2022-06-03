package square

import "testing"

func TestArea(t *testing.T) {
	square := Square{start: Point{x: 3, y: 5}, a: 3}

	want := uint(9)
	got := square.Area()

	if got != want {
		t.Errorf("TestArea got: %v, want: %v", got, want)
	}
}

func TestPermiter(t *testing.T) {
	square := Square{start: Point{x: 8, y: 9}, a: 5}

	want := uint(20)
	got := square.Perimeter()

	if got != want {
		t.Errorf("TestPermiter got: %v, want: %v", got, want)
	}
}

func TestEnd(t *testing.T) {
	square := Square{start: Point{x: 8, y: 9}, a: 5}
	want := Point{x: 13, y: 14}
	got := square.End()

	if got != want {
		t.Errorf("TestEnd got: %v, want: %v", got, want)
	}
}
