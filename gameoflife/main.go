package main

import (
	"context"
	"flag"
	"image"
	"image/color"
	"strings"
	"time"

	"github.com/rivo/tview"
	"k8s.io/klog/v2"
)

var (
	dx = flag.Int("dx", 32, "Horizontal size of the grid")
	dy = flag.Int("dy", 32, "Vertical size of the grid")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	a := newArena(*dx, *dy)

	// A blinker
	// for x := 3; x < 6; x++ {
	// 	a.current[2][x] = true
	// }

	a.current[18][11] = true
	a.current[18][13] = true
	a.current[17][13] = true
	a.current[14][15] = true
	a.current[15][15] = true
	a.current[16][15] = true
	a.current[13][17] = true
	a.current[14][17] = true
	a.current[15][17] = true
	a.current[14][18] = true

	app := tview.NewApplication()
	image := tview.NewImage()

	go func() {
		t := time.NewTicker(200 * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				klog.Info("Evolve function quitting")
				return
			case <-t.C:
			}
			a.evolve()
			i := a.Image()
			image.SetImage(i)
			app.Draw()
		}
	}()

	app.SetRoot(image, true).Run()
}

func newArena(dx, dy int) arena {
	// Allocate all cells in one go so they are contiguous
	current := make([][]bool, dy)
	currentCells := make([]bool, dx*dy)
	next := make([][]bool, dy)
	nextCells := make([]bool, dx*dy)
	for y := 0; y < dy; y++ {
		current[y] = currentCells[y*dx : (y+1)*dx]
		next[y] = nextCells[y*dx : (y+1)*dx]
	}
	return arena{
		dx:      dx,
		dy:      dy,
		current: current,
		next:    next,
	}
}

type arena struct {
	dx, dy        int
	current, next [][]bool
}

func (a *arena) evolve() {
	for y := 0; y < len(a.current); y++ {
		lastRow := a.current[(y-1+a.dy)%a.dy]
		row := a.current[y]
		nextRow := a.current[(y+1)%a.dy]
		for x := 0; x < len(row); x++ {
			xPrev := (x - 1 + a.dx) % a.dx
			xNext := (x + 1) % a.dx
			neighbours := 0
			if lastRow[xPrev] {
				neighbours++
			}
			if lastRow[x] {
				neighbours++
			}
			if lastRow[xNext] {
				neighbours++
			}
			if row[xPrev] {
				neighbours++
			}
			if row[xNext] {
				neighbours++
			}
			if nextRow[xPrev] {
				neighbours++
			}
			if nextRow[x] {
				neighbours++
			}
			if nextRow[xNext] {
				neighbours++
			}
			live := row[x]
			if !live {
				a.next[y][x] = neighbours == 3
				continue
			}
			a.next[y][x] = neighbours == 2 || neighbours == 3
		}
	}
	a.current, a.next = a.next, a.current
}

func (a *arena) String() string {
	sb := strings.Builder{}
	for y := 0; y < *dy; y++ {
		for x := 0; x < *dx; x++ {
			cell := a.current[y][x]
			if cell {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (a *arena) Image() image.Image {
	// TODO(mhutchinson): Reuse images and write to Pix directly
	i := image.NewNRGBA(image.Rect(0, 0, a.dx, a.dy))
	for y := 0; y < *dy; y++ {
		for x := 0; x < *dx; x++ {
			cell := a.current[y][x]
			c := color.White
			if cell {
				c = color.Black
			}
			i.Set(x, y, c)
		}
	}
	return i
}
