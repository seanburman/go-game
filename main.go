package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/seanburman/gogame/Game"
	"github.com/seanburman/gogame/characters"
	"github.com/seanburman/gogame/constants"
	"github.com/seanburman/gogame/graphics"
)

var r *graphics.Registry

func init() {
	r = &graphics.Registry{}
	r.Register([]graphics.Graphic{
		new(characters.Player),
		new(characters.Female),
	})
}

func main() {
	ebiten.SetWindowSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")

	g := &game.Game{
		Graphics: r,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
