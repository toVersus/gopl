// ex3.7 visualizes how many iterations it takes to find complex roots of a
// quartic equation using Newton's method, using different colors for
// different roots.
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

type Func func(complex128) complex128

var colorPool = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

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
			// Image point (px, py) represents complex value z.
			img.Set(px, py, z4(z))
		}
	}
	f, err := os.Create("newton.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file: %s", err)
	}
	png.Encode(f, img) // NOTE: ignoring errors
}

func z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func newton(z complex128, f, fPrime Func) color.Color {
	const iterations = 200
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(colorPool) == 0 {
					panic("no colors left")
				}
				chosenColors[root], colorPool = colorPool[0], colorPool[1:]
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	pow := math.Pow10(digits)
	return math.Floor(f*pow+.5) / pow
}
