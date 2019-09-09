package main

import (
	"github.com/fogleman/gg"
	"github.com/mhutchinson/goplayground/fractals/mandelbrot"
)

const (
	width  = 640
	height = 640
)

func main() {
	drawMandlebrot(-2+2i, 2-2i)
}

func drawMandlebrot(topLeft, bottomRight complex128) {
	mandlebrot := mandelbrot.NewQuadraticCalculator(500)
	xInc := (real(bottomRight) - real(topLeft)) / width
	yInc := (imag(bottomRight) - imag(topLeft)) / height
	dc := gg.NewContext(width, height)
	realComp, imagComp := real(topLeft), imag(topLeft)
	for x := 0; x < dc.Width(); x++ {
		imagComp = imag(topLeft)
		for y := 0; y < dc.Height(); y++ {
			point := complex(realComp, imagComp)
			score := mandlebrot.Score(point)
			if x == y {
				println(x, y, point, score)
			}
			dc.SetRGB(score, score, score)
			dc.SetPixel(x, y)
			imagComp += yInc
		}
		realComp += xInc
	}
	dc.SavePNG("mandelbrot.png")
}
