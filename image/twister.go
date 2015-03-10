// +build gm
package image

import (
    m "gopkgs.com/magick.v1"
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

func Scale(image *m.Image, x integer,y integer) *m.Image {
	scaled, _ := image.Scale(x, y)
	return scaled
}
