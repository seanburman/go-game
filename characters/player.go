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
	// pos.Update = *p.Update()
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

func (p *Player) Draw(screen *ebiten.Image) {
	p.Character.Draw(screen)
}

func (p *Player) Position() graphics.Position {
	return p.Character.Pos
}

func (p *Player) Move() {
	pos := &p.Character.Pos
	// Left
	if ok := ebiten.IsKeyPressed(ebiten.KeyA); ok {
		// Set direction
		pos.Direction = "LEFT"
		pos.Vector2 = [2]int{-1, 0}
		// Update position
		if pos.X > 0 {
			pos.X -= p.Character.Pos.Speed
		}
	}
	// Right
	if ok := ebiten.IsKeyPressed(ebiten.KeyD); ok {
		// Set direction
		pos.Direction = "RIGHT"
		pos.Vector2 = [2]int{1, 0}
		// Update position
		if pos.X < constants.SCREEN_WIDTH-float64(pos.Max.X) {
			pos.X += pos.Speed
		}
	}
	// Up
	if ok := ebiten.IsKeyPressed(ebiten.KeyW); ok {
		// Set direction
		pos.Direction = "UP"
		pos.Vector2 = [2]int{0, -1}
		// Update position
		if pos.Y > 0 {
			pos.Y -= pos.Speed
		}
	}
	// Down
	if ok := ebiten.IsKeyPressed(ebiten.KeyS); ok {
		// Set direction
		pos.Direction = "DOWN"
		pos.Vector2 = [2]int{0, 1}
		// Update position
		if pos.Y < constants.SCREEN_HEIGHT-float64(pos.Max.Y) {
			pos.Y += pos.Speed
		}
	}
}
