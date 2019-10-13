package main

import (
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/fogleman/gg"
	"github.com/mhutchinson/goplayground/fractals/mandelbrot"
)

const (
	w, h   = 512, 512
	fw, fh = float64(w), float64(h)
	width  = 1024
	height = 1024
)

func main() {
	pixelgl.Run(run)
	// treeFun()
	// drawMandlebrot(-2.25+1.5i, 0.75-1.5i, "mandelbrot.png")
	// drawMandlebrot(-0.74540+0.11260i, -0.74535+0.11255i, "hardzoom.png")
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, w, h),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	calculator := mandelbrot.NewQuadraticCalculator(500)
	canvas := pixelgl.NewCanvas(win.Bounds())

	const quotient = 50
	destTL, destBR := -0.74540+0.11260i, -0.74535+0.11255i

	topLeft, bottomRight := -2.25+1.5i, 0.75-1.5i
	for !win.Closed() {
		buffer := image.NewRGBA(image.Rect(0, 0, w, h))

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		t := mandelbrot.NewTile(topLeft, bottomRight, h, w)
		grid := t.Calculate(calculator)

		for x := 0; x < len(grid); x++ {
			col := grid[x]
			for y := 0; y < len(col); y++ {
				score := grid[x][y]
				colorInt := uint8(255 * score)
				buffer.Set(x, y, color.RGBA{colorInt, colorInt, colorInt, 1.0})
			}
		}

		win.Clear(color.RGBA{0, 0, 0, 255})
		canvas.SetPixels(buffer.Pix)
		canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()

		topLeft = topLeft - (topLeft-destTL)/quotient
		bottomRight = bottomRight - (bottomRight-destBR)/quotient
	}
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
