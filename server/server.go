package server

import (
    twister "brutalbits.com/pixeltwister/image"
    "net/http"
)

func Start() {
    http.HandleFunc("/load", load)
    http.ListenAndServe(":8080", nil)
}

func load(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "image/png")

    artFile := twister.LoadImage("data/base.jpg")
    twister.EncodePng(response, artFile)
}
