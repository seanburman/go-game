package graphics

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewImageFromFile(file string) *ebiten.Image {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func NewImagesFromFiles(f map[string]string) map[string]*ebiten.Image {
	images := map[string]*ebiten.Image{}

	for k, v := range f {
		img := NewImageFromFile(v)
		images[k] = ebiten.NewImageFromImage(img)
	}
	return images
}
