package main

import (
	"flag"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mhutchinson/goplayground/fractals/mandelbrot"
)

var (
	iters = flag.Int("i", 200, "Number of iterations per point to determine set membership. Higher values are more accurate but slower.")
)

const (
	w, h     = 800, 800
	width    = 1024
	height   = 1024
	quotient = 50
)

type Game struct {
	calculator  *mandelbrot.Calculator
	topLeft     complex128
	bottomRight complex128
	destTL      complex128
	destBR      complex128
	pixels      []byte
}

func NewGame() *Game {
	return &Game{
		calculator:  mandelbrot.NewQuadraticCalculator(*iters),
		topLeft:     -2.25 + 1.5i,
		bottomRight: 0.75 - 1.5i,
		destTL:      -0.74540 + 0.11260i,
		destBR:      -0.74535 + 0.11255i,
		pixels:      make([]byte, w*h*4),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	t := mandelbrot.NewTile(g.topLeft, g.bottomRight, h/2, w/2)
	grid := t.Calculate(g.calculator)

	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col); y++ {
			score := grid[x][y]
			colorInt := uint8(255 * score)
			g.set2x2(x, y, colorInt)
		}
	}

	g.topLeft = g.topLeft - (g.topLeft-g.destTL)/quotient
	g.bottomRight = g.bottomRight - (g.bottomRight-g.destBR)/quotient

	return nil
}

func (g *Game) set2x2(x, y int, colorVal uint8) {
	setPixel := func(px, py int) {
		idx := (py*w + px) * 4
		g.pixels[idx] = colorVal   // R
		g.pixels[idx+1] = colorVal // G
		g.pixels[idx+2] = colorVal // B
		g.pixels[idx+3] = 255      // A
	}
	setPixel(2*x, 2*y)
	setPixel(2*x+1, 2*y)
	setPixel(2*x, 2*y+1)
	setPixel(2*x+1, 2*y+1)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func main() {
	flag.Parse()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Mandelbrot Zoom")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	// treeFun()
	// drawMandlebrot(-2.25+1.5i, 0.75-1.5i, "mandelbrot.png")
	// drawMandlebrot(-0.74540+0.11260i, -0.74535+0.11255i, "hardzoom.png")
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
