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
	treeFun()
	drawMandlebrot(-2.25+1.5i, 0.75-1.5i, "mandelbrot.png")
	drawMandlebrot(-0.74540+0.11260i, -0.74535+0.11255i, "hardzoom.png")
}

func drawMandlebrot(topLeft, bottomRight complex128, file string) {
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
	dc.SavePNG(file)
}

// TODO(mhutchinson): refactor this into its own class
func treeFun() {
	dc := gg.NewContext(width, height)
	dc.Translate(width/2, height)

	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(5)
	drawTree(dc, height/4)
	dc.Stroke()
	dc.SavePNG("treefun.png")
}

const angle float64 = 45
const ratio float64 = 0.7

func drawTree(dc *gg.Context, len float64) {
	if len < 5 {
		return
	}
	dc.DrawLine(0, 0, 0, -len)
	dc.Translate(0, -len)
	dc.Push()
	dc.Rotate(angle)
	drawTree(dc, len*ratio)
	dc.Pop()
	dc.Push()
	dc.Rotate(-angle)
	drawTree(dc, len*ratio)
	dc.Pop()
}
