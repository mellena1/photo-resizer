package photos

import (
	"image"
	"os"

	"github.com/bmatcuk/doublestar"

	log "github.com/sirupsen/logrus"
)

// FileImage is a Image with its location/name
type FileImage struct {
	Filename string      `json:"file"`
	Img      image.Image `json:"-"`
}

// LoadImage loads up FileImage as an image with only the filename
func (f *FileImage) LoadImage() error {
	img, err := Open(f.Filename)
	if err != nil {
		log.Warningf("%s couldn't be opened as an image: %v", f.Filename, err)
		return err
	}
	f.Img = img
	return nil
}

// FindImagesFromGlob get all FileImage structs from a file glob
func FindImagesFromGlob(glob string) ([]*FileImage, error) {
	matches, err := doublestar.Glob(glob)
	if err != nil {
		return nil, err
	}
	if matches == nil {
		return []*FileImage{}, nil
	}

	images := []*FileImage{}
	for _, match := range matches {
		if info, _ := os.Stat(match); info.IsDir() {
			continue
		}
		img := &FileImage{Filename: match}
		err := img.LoadImage()
		if err == nil {
			images = append(images, img)
		}
	}
	return images, nil
}
