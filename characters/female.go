package characters

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanburman/gogame/constants"
	"github.com/seanburman/gogame/graphics"
)

// IMPLEMENTATIONS OF CHARACTER
//

type Female struct {
	Character *Character
}

func (p *Female) Create(r *graphics.Registry) {
	pos := graphics.Position{
		X:         200,
		Y:         200,
		Direction: "UP",
		Speed:     constants.SPEED,
	}
	// pos.Update = *p.Update()
	imgPaths := map[string]string{
		"UP":    "assets/img/sprites/female/up.png",
		"DOWN":  "assets/img/sprites/female/down.png",
		"LEFT":  "assets/img/sprites/female/left.png",
		"RIGHT": "assets/img/sprites/female/right.png",
	}
	c := Character{}
	c.Create(pos, imgPaths)
	p.Character = &c
}

func (p *Female) Update() error {
	p.Character.Update()
	return nil
}

func (p *Female) Draw(screen *ebiten.Image, r *graphics.Registry) {
	p.Character.Draw(screen, r)
}

func (p *Female) Position() graphics.Position {
	return p.Character.Pos
}
