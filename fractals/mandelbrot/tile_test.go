package mandelbrot

import (
	"testing"
)

func BenchmarkCalculate(b *testing.B) {
	t := NewTile(-2+2i, 2-2i, 500, 500)
	calc := NewQuadraticCalculator(100)

	for i := 0; i < b.N; i++ {
		t.Calculate(calc)
	}
}
