package server

import (
	twister "brutalbits.com/pixeltwister/image"
	"net/http"
)

func Start() {
	http.HandleFunc("/load", load)
	http.HandleFunc("/thumb400", thumb400)
	http.HandleFunc("/sharpThumb400", sharpThumb400)
	http.ListenAndServe(":8080", nil)
}

func load(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "image/png")
	result := twister.LoadImage("data/base.jpg")
	twister.EncodePng(response, result)
}

func thumb400(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "image/png")
	image := twister.LoadImage("data/base.jpg")
	result := twister.Scale(image, 400, 400)
	twister.EncodePng(response, result)
}

func sharpThumb400(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "image/png")
	image := twister.LoadImage("data/base.jpg")
	scaled := twister.Scale(image, 400, 400)
	result := twister.UnsharpMask(scaled, 0.0, 1.0, 0.9, 0.1)
	twister.EncodePng(response, result)
}
