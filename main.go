package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type vec2 struct {
	X float32
	Y float32
}

func Vec2(x, y float32) vec2 {
	ret := vec2{}
	ret.X = x
	ret.Y = y
	return ret
}

type Game struct {
	stroke []vec2
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.stroke = append(g.stroke, Vec2(float32(x), float32(y)))

	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.stroke = []vec2{}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	out := ""
	out += fmt.Sprintf("%.2f\n", ebiten.ActualFPS())
	out += fmt.Sprintf("pressing: %t\n", ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
	ebitenutil.DebugPrint(screen, out)

	for i := range g.stroke {
		if i == 0 {
			continue
		}

		from := g.stroke[i-1]
		to := g.stroke[i]
		vector.StrokeLine(screen, from.X, from.Y, to.X, to.Y, 1, color.White, false)
	}

	if 0 < len(g.stroke) {
		end := g.stroke[len(g.stroke)-1]

		vector.StrokeCircle(screen, end.X, end.Y, 3, 1, color.White, false)
	}
}

func (g *Game) Layout(outerWidth, outerHeight int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("demo")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
