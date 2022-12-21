package imutil

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/jo-m/trainbot/pkg/cqoi"
)

// Load tries to load an image from a file.
func Load(path string) (image.Image, error) {
	if strings.HasSuffix(path, ".qoi") {
		return cqoi.Load(path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Dump will dump an image to a file.
// Format is determined by file ending, PNG, JPEG and QOI are supported.
func Dump(path string, img image.Image) error {
	if strings.HasSuffix(path, ".qoi") {
		return cqoi.Dump(path, img)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if strings.HasSuffix(path, ".png") {
		return png.Encode(f, img)
	}

	if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		return jpeg.Encode(f, img, &jpeg.Options{Quality: 98})
	}

	return errors.New("unknown image suffix")
}

// Sub tries to call SubImage on the given image.
func Sub(img image.Image, r image.Rectangle) (image.Image, error) {
	iface, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	})

	if !ok {
		return nil, errors.New("img does not implement SubImage()")
	}

	return iface.SubImage(r), nil
}
