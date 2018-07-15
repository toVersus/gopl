// ex3.6 emits a supersampled image of the Mandelbrot fractal.
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
		xpixel                 = (xmax - xmin) / width
		ypixel                 = (ymax - ymin) / height
	)

	xoffset := []float64{-xpixel, xpixel}
	yoffset := []float64{-ypixel, ypixel}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			subPixels := make([]color.Color, 0)
			for i := range xoffset {
				for j := range yoffset {
					z := complex(x+xoffset[i], y+yoffset[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			// Image point (px, py) represents vomplex value z.
			img.Set(px, py, avg(subPixels))
		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file: %s", err)
	}
	png.Encode(f, img) // NOTE: ignoring errors
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	size := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(size))
		g += uint16(g_ / uint32(size))
		b += uint16(b_ / uint32(size))
		a += uint16(a_ / uint32(size))
	}
	return color.RGBA64{r, g, b, a}
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
