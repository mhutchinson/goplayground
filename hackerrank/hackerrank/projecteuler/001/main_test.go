package main

import (
	"testing"
)

func TestMultiplesSum(t *testing.T) {
	tables := []struct {
		n    int32
		want int64
	}{
		{3, 0},
		{4, 3},
		{5, 3},
		{6, 8},
		{7, 14},
		{10, 23},
		{16, 60},
		{31, 225},
		{46, 495},
		{61, 870},
	}
	for _, c := range tables {
		got := MultiplesSum(c.n)
		if got != c.want {
			t.Errorf("MultiplesSum(%d): got %d, want %d", c.n, got, c.want)
		}
	}
}

func BenchmarkMultiplesSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplesSum(5000)
	}
}
