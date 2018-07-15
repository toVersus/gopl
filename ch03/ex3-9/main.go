// ex3.9 serves images of fractals over http.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	const width, height = 1024, 1024
	var (
		xmin, xmax float64 = -2, 2
		ymin, ymax float64 = -2, 2
		zoom       float64 = 1
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		xmins, ok := r.URL.Query()["xmin"]
		if ok {
			xmin, _ = strconv.ParseFloat(xmins[0], 64)
		}
		xmaxs, ok := r.URL.Query()["xmax"]
		if ok {
			xmax, _ = strconv.ParseFloat(xmaxs[0], 64)
		}
		ymins, ok := r.URL.Query()["ymin"]
		if ok {
			ymin, _ = strconv.ParseFloat(ymins[0], 64)
		}
		ymaxs, ok := r.URL.Query()["ymax"]
		if ok {
			ymax, _ = strconv.ParseFloat(ymaxs[0], 64)
		}
		zooms, ok := r.URL.Query()["zoom"]
		if ok {
			zoom, _ = strconv.ParseFloat(zooms[0], 64)
		}

		if xmax <= xmin || ymax <= ymin {
			http.Error(w, fmt.Sprintf("min coordinate greater than max"), http.StatusBadRequest)
			return
		}

		xlen := xmax - xmin
		xmid := xmin + xlen/2
		xmin = xmid - xlen/2/zoom
		xmax = xmid + xlen/2/zoom
		ylen := ymax - ymin
		ymid := ymin + ylen/2
		ymin = ymid - ylen/2/zoom
		ymax = ymid + ylen/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value
				img.Set(px, py, mandelbrot(z))
			}
		}
		err := png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
