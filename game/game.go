package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanburman/gogame/constants"
	"github.com/seanburman/gogame/graphics"
)

type Game struct {
	Graphics *graphics.Registry
}

func (g *Game) Update() error {
	g.Graphics.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0})
	g.Graphics.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT
}
