package photos

import (
	"image"

	"github.com/disintegration/imaging"
)

// Open an image
func Open(filename string) (image.Image, error) {
	return imaging.Open(filename)
}

// Resize an image
func Resize(img *FileImage, savename string, width, height int) {
	dest := imaging.Resize(img.Img, width, height, imaging.Lanczos)
	imaging.Save(dest, savename)
}
