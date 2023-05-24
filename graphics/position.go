package graphics

import "github.com/seanburman/gogame/constants"

type (
	Position struct {
		X         float64
		Y         float64
		Vector2   [2]int
		Direction string
		Speed     float64
		Max       struct {
			X int
			Y int
		}
		Mover Mover
	}
)

type Mover interface {
	Move()
}

func (p *Position) Move(r *Registry) {
	if !p.CollidesWithSlice(r.Graphics) && p.Mover != nil {
		p.Mover.Move()
	}
}

func (p *Position) Left() {
	// Set direction
	p.Direction = "LEFT"
	// Set velocity
	p.Vector2 = [2]int{-1 * int(p.Speed), 0}
	// Update position
	if p.X > 0 {
		p.X -= p.Speed
	}
}

func (p *Position) Right() {
	// Set direction
	p.Direction = "RIGHT"
	// Set velocity
	p.Vector2 = [2]int{int(p.Speed), 0}
	// Update position
	if p.X < constants.SCREEN_WIDTH-float64(p.Max.X) {
		p.X += p.Speed
	}
}

func (p *Position) Up() {
	// Set direction
	p.Direction = "UP"
	// Set velocity
	p.Vector2 = [2]int{0, -1 * int(p.Speed)}
	// Update position
	if p.Y > 0 {
		p.Y -= p.Speed
	}
}

func (p *Position) Down() {
	// Set direction
	p.Direction = "DOWN"
	// Set velocity
	p.Vector2 = [2]int{0, int(p.Speed)}
	// Update position
	if p.Y < constants.SCREEN_HEIGHT-float64(p.Max.Y) {
		p.Y += p.Speed
	}
}

func (p *Position) Collides(p2 Position) bool {
	// aabb algorithm for detecting collision of vertices
	// a.x < b.x + b.width
	// && a.x + a.width > b.x
	// && a.y < b.y + b.height
	// && a.y + a.height > b.y
	return (p.X < p2.X+float64(p2.Max.X) && p.X+float64(p.Max.X) > p2.X && p.Y < p2.Y+float64(p2.Max.Y) && p.Y+float64(p.Max.Y) > p2.Y)
}

// Returns whether this *Position intersects with others in Registry Graphics
func (p *Position) CollidesWithSlice(gs []Graphic) bool {
	for _, g := range gs {
		if g.Position() == *p {
			continue
		}
		if p.Collides(g.Position()) {
			return true
		}
	}
	return false
}
