// ex3.5 emits a full-color PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents vomplex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file: %s", err)
	}
	png.Encode(f, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			logScale := math.Log(float64(n)) / math.Log(float64(iterations))
			return color.RGBA{100, 255 - uint8(logScale*255), 100, 255}
		}
	}
	return color.Black
}
