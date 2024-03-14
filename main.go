package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	w, h := 1920, 1080
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Shader Issue")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	options := &ebiten.DrawRectShaderOptions{}
	options.GeoM.Translate(100, 100)
	screen.DrawRectShader(100, 100, Shader, options)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return 1920, 1080
}
