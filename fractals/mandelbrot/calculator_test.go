package mandelbrot

import (
	"fmt"
	"math"
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		input complex128
		want  int
	}{
		{-1 + 0i, 100},
		{1 + 0i, 0},
		{0 + 1i, 100},
		{0 - 1i, 100},
	}
	calc := NewQuadraticCalculator(500)
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			got := int(math.Round(100 * calc.Score(c.input)))
			if got != c.want {
				t.Errorf("Score(%v): got %d, want %d", c.input, got, c.want)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	inputs := []complex128{
		-1 + 0i,
		1 + 0i,
		0 + 1i,
		0 - 1i,
	}
	calc := NewQuadraticCalculator(500)
	for _, input := range inputs {
		b.Run(fmt.Sprintf("%v", input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				calc.Score(input)
			}
		})
	}
}
