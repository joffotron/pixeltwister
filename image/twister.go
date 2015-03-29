// +build gm
package image

import (
	m "brutalbits.com/magick"
	"io"
)

func LoadImage(filename string) *m.Image {
	image, _ := m.DecodeFile(filename)
	return image
}

func EncodePng(output io.Writer, image *m.Image) {
	info := m.NewInfo()
	info.SetFormat("PNG")
	image.Encode(output, info)
}

func Scale(image *m.Image, x int, y int) *m.Image {
	scaled, _ := image.Scale(x, y)
	return scaled
}

func UnsharpMask(image *m.Image, radius float64, sigma float64, amount float64, threshold float64) *m.Image {
	sharpened, _ := image.UnsharpMask(radius, sigma, amount, threshold)
	return sharpened
}

func Compose(baseImage *m.Image, overlay *m.Image, x int, y int) *m.Image {
	baseImage.Composite(m.CompositeOver, overlay, x, y)
	return baseImage
}
