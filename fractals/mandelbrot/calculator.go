package mandelbrot

import (
	"math/cmplx"
)

type algorithm func(complex128, complex128) complex128

// Calculator has methods for performing core mandelbrot operations.
// This class is thread-safe.
type Calculator struct {
	maxIterations int
	alg           algorithm
}

// Score returns a value between 0 and 1 for the probability that p is
// in the Mandelbrot set.
func (c *Calculator) Score(p complex128) float64 {
	z := p
	var i int
	for ; i < c.maxIterations; i++ {
		// TODO(mhutchinson): potentially look for cycles and break earlier.
		// The tightness of the cycle could be interesting data in the return type.
		z = c.alg(z, p)
		if cmplx.Abs(z) > 2 {
			break
		}
	}
	return float64(i) / float64(c.maxIterations)
}

// NewQuadraticCalculator returns a Calculator which classifies members for the classic mandelbrot set.
func NewQuadraticCalculator(maxIterations int) *Calculator {
	var alg algorithm = func(z, p complex128) complex128 { return z*z + p }
	return &Calculator{
		maxIterations: maxIterations,
		alg:           alg,
	}
}
