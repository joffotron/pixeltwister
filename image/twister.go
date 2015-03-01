package image

import (
    magick "gopkgs.com/magick.v1"
    "io"
)

func LoadImage(filename string) *magick.Image {
    image, _ := magick.DecodeFile(filename)
    return image
}

func EncodePng(output io.Writer, image *magick.Image) {
    info := magick.NewInfo()
    info.SetFormat("PNG")
    image.Encode(output, info)
}
