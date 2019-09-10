package main

import (
	"github.com/fogleman/gg"
	"github.com/mhutchinson/goplayground/fractals/mandelbrot"
)

const (
	width  = 1024
	height = 1024
)

func main() {
	drawMandlebrot(-2.25+1.5i, 0.75-1.5i)
}

func drawMandlebrot(topLeft, bottomRight complex128) {
	calculator := mandelbrot.NewQuadraticCalculator(500)
	t := mandelbrot.NewTile(topLeft, bottomRight, height, width)
	grid := t.Calculate(calculator)
	dc := gg.NewContext(width, height)
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col); y++ {
			score := grid[x][y]
			dc.SetRGB(score, score, score)
			dc.SetPixel(x, y)
		}
	}
	dc.SavePNG("mandelbrot.png")
}
