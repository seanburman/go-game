package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Graphic interface {
	Create(gr *Registry)
	Draw(screen *ebiten.Image)
	Update() error
	Position() Position
}

type Registry struct {
	Graphics []Graphic
}

func (r *Registry) Register(gr []Graphic) {
	r.Graphics = gr
	r.Create()
}

func (r *Registry) Create() {
	for _, gr := range r.Graphics {
		gr.Create(r)
	}
}

func (r *Registry) Update() {
	for _, gr := range r.Graphics {
		gr.Update()
	}
}

func (r *Registry) Draw(screen *ebiten.Image) {
	for _, gr := range r.Graphics {
		gr.Draw(screen)
	}
}
