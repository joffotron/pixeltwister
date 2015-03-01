package image

import (
    "image"
    _ "image/jpeg"
    "os"
    "brutalbits.com/pixeltwister/image"
)

func LoadImage(filename string) image.Image {
    file, _ := os.Open(filename)
    defer file.Close()

    image, _, _ := image.Decode(file)
    return image
}
