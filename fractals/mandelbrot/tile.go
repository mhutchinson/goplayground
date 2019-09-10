package mandelbrot

import (
	"sync"
)

// Tile represents a rectangular area in the complex number plane.
type Tile struct {
	topLeft, bottomRight complex128
	rows, cols           int
}

// NewTile makes a new Tile for the specified area, with the given resolution.
func NewTile(topLeft, bottomRight complex128, rows, cols int) Tile {
	return Tile{
		topLeft:     topLeft,
		bottomRight: bottomRight,
		rows:        rows,
		cols:        cols,
	}
}

// Calculate returns the result of scoring each cell in this tile.
// (0,0) is the top left, (cols-1, rows-1) is the bottom right.
func (t *Tile) Calculate(c *Calculator) [][]float64 {
	xInc := (real(t.bottomRight) - real(t.topLeft)) / float64(t.cols)
	yInc := (imag(t.bottomRight) - imag(t.topLeft)) / float64(t.rows)
	realComp := real(t.topLeft)

	var wg sync.WaitGroup
	wg.Add(t.cols)
	r := make([][]float64, t.cols)
	for x := 0; x < t.cols; x++ {
		r[x] = make([]float64, t.rows)
		go func(x int, realComp float64) {
			defer wg.Done()
			imagComp := imag(t.topLeft)
			for y := 0; y < t.rows; y++ {
				point := complex(realComp, imagComp)
				score := c.Score(point)
				r[x][y] = score
				imagComp += yInc
			}
		}(x, realComp)
		realComp += xInc
	}
	wg.Wait()
	return r
}
