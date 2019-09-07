package fib

import "testing"

func RunTests(t *testing.T, fib func(int) int) {
	tables := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}
	for _, table := range tables {
		got := fib(table.n)
		if got != table.want {
			t.Errorf("fib(%d): got %d, want %d", table.n, got, table.want)
		}
	}
}
