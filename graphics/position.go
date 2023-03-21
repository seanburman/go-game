package graphics

import (
	"fmt"
)

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

func (p *Position) Move() {
	if !p.Intersects([]Position{}) && p.Mover != nil {
		p.Mover.Move()
	}
}

// Returns whether this *Position intersects with any in a slice of Position
func (p *Position) Intersects(pslice []Position) bool {
	for k, v := range pslice {
		fmt.Println(k, v)
	}
	return false
}
