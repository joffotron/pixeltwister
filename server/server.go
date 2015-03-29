package server

import (
	m "brutalbits.com/magick"
	twister "brutalbits.com/pixeltwister/image"
	"net/http"
)

func Start() {
	http.HandleFunc("/load", load)
	http.HandleFunc("/thumb400", thumb400)
	http.HandleFunc("/sharpThumb400", sharpThumb400)
	http.HandleFunc("/allTheThings", allTheThings)
	http.ListenAndServe(":8080", nil)
}

func load(response http.ResponseWriter, request *http.Request) {
	setImageHeader(response)
	result := twister.LoadImage("data/base.jpg")
	twister.EncodePng(response, result)
}

func thumb400(response http.ResponseWriter, request *http.Request) {
	setImageHeader(response)
	result := twister.Scale(twister.LoadImage("data/base.jpg"), 400, 400)
	twister.EncodePng(response, result)
}

func sharpThumb400(response http.ResponseWriter, request *http.Request) {
	setImageHeader(response)
	result := someOfTheThings(twister.LoadImage("data/base.jpg"))
	twister.EncodePng(response, result)
}

func allTheThings(response http.ResponseWriter, request *http.Request) {
	setImageHeader(response)
	baseImage := twister.LoadImage("data/base.jpg")
	some := someOfTheThings(baseImage)
	thumb := twister.Scale(baseImage, 200, 150)
	result := twister.Compose(some, thumb, 40, 40)
	twister.EncodePng(response, result)
}

// ================++++++++++++==============+++++++++++++=========

func someOfTheThings(image *m.Image) *m.Image {
	scaled := twister.Scale(image, 800, 600)
	sharpened := twister.UnsharpMask(scaled, 0.0, 1.0, 0.9, 0.1)
	return sharpened
}

func setImageHeader(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "image/png")
}
