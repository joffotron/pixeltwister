package server

import (
    twister "brutalbits.com/pixeltwister/image"
    "net/http"
)

func Start() {
    http.HandleFunc("/load", load)
    http.HandleFunc("/thumb400", thumb400)
    http.ListenAndServe(":8080", nil)
}

func load(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "image/png")

    image := twister.LoadImage("data/base.jpg")
    twister.EncodePng(response, image)
}

func thumb400(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "image/png")
    image := twister.LoadImage("data/base.jpg")
    scaled, _ := image.Scale(400, 400)
    twister.EncodePng(response, scaled)
}
