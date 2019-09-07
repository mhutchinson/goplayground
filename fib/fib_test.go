package fib

import "testing"

var impls = []struct {
	name string
	fib  func(int) int
}{
	{"recursive", RecursiveFib},
	{"parallelRecursive", ParallelRecursiveFib},
	{"iterative", IterativeFib},
	{"iterativeGen", IterativeGenFib},
}

func TestFib(t *testing.T) {
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
	for _, impl := range impls {
		t.Run(impl.name, func(t *testing.T) {
			for _, table := range tables {
				got := impl.fib(table.n)
				if got != table.want {
					t.Errorf("fib(%d): got %d, want %d", table.n, got, table.want)
				}
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				impl.fib(30)
			}
		})
	}
}
