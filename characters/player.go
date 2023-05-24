package characters

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanburman/gogame/constants"
	"github.com/seanburman/gogame/graphics"
)

// IMPLEMENTATIONS OF CHARACTER
//

type Player struct {
	Character *Character
}

func (p *Player) Create(r *graphics.Registry) {
	pos := graphics.Position{
		X:         50,
		Y:         constants.WINDOW_HEIGHT / 2,
		Direction: "DOWN",
		Speed:     constants.SPEED,
		Mover:     p,
	}
	imgPaths := map[string]string{
		"UP":    "assets/img/sprites/player/up.png",
		"DOWN":  "assets/img/sprites/player/down.png",
		"LEFT":  "assets/img/sprites/player/left.png",
		"RIGHT": "assets/img/sprites/player/right.png",
	}
	c := Character{}
	c.Create(pos, imgPaths)
	p.Character = &c
}

func (p *Player) Update() error {
	p.Character.Update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image, r *graphics.Registry) {
	p.Character.Draw(screen, r)
}

func (p *Player) Position() graphics.Position {
	return p.Character.Pos
}

func (p *Player) Move() {
	pos := &p.Character.Pos
	// Left
	if ok := ebiten.IsKeyPressed(ebiten.KeyA); ok {
		pos.Left()
	}
	// Right
	if ok := ebiten.IsKeyPressed(ebiten.KeyD); ok {
		pos.Right()
	}
	// Up
	if ok := ebiten.IsKeyPressed(ebiten.KeyW); ok {
		pos.Up()
	}
	// Down
	if ok := ebiten.IsKeyPressed(ebiten.KeyS); ok {
		pos.Down()
	}
}
