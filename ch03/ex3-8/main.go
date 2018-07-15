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
	"math/big"
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
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigFloat(z))
		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file: %s", err)
	}
	png.Encode(f, img) // NOTE: ignoring errors
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := 0; n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			logScale := math.Log(float64(n) / math.Log(float64(iterations)))
			return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + complex128(z)
		if cmplx.Abs(complex128(v)) > 2 {
			if n > 50 { // dark red
				return color.RGBA{100, 0, 0, 255}
			}
			logScale := math.Log(float64(n) / math.Log(float64(iterations)))
			return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := 0; i < iterations; i++ {
		// (r + i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			logScale := math.Log(float64(i)) / math.Log(float64(iterations))
			return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
		}
	}
	return color.Black
}
