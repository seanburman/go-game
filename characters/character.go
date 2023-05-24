package characters

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanburman/gogame/graphics"
)

type (
	Character struct {
		Img  *ebiten.Image
		Imgs map[string]*ebiten.Image
		Pos  graphics.Position
	}
)

func (c *Character) Create(pos graphics.Position, imgPaths map[string]string) {
	c.Pos = pos
	c.Pos.Vector2 = [2]int{0, 0}
	c.GetImages(imgPaths)
}

func (c *Character) Draw(screen *ebiten.Image, r *graphics.Registry) {
	// check current direction
	d := c.Pos.Direction
	// update position from user entry
	c.Pos.Move(r)
	// check if direction has changed since update
	if d != c.Pos.Direction {
		// change image to match direction
		c.ChangeImage(c.Pos.Direction)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.Pos.X, c.Pos.Y)
	screen.DrawImage(c.Img, op)
}

func (c *Character) Update() {
}

func (c *Character) GetImages(imgPaths map[string]string) {
	c.Imgs = graphics.NewImagesFromFiles(imgPaths)
	c.Img = c.Imgs[c.Pos.Direction]
}

func (c *Character) ChangeImage(ImgsKey string) {
	// update Img
	c.Img = c.Imgs[ImgsKey]
	// update bounds
	c.Pos.Max.X = c.Img.Bounds().Max.X
	c.Pos.Max.Y = c.Img.Bounds().Max.Y
}
