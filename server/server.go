package server

import (
    twister "brutalbits.com/pixeltwister/image"
    "image/png"
    "net/http"
)

func Start() {
    http.HandleFunc("/load", load)
    http.ListenAndServe(":8080", nil)
}

func load(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "image/png")
    //func DetectContentType(data []byte) string

    artFile := twister.LoadImage("data/base.jpg")
    png.Encode(response, artFile)
}
