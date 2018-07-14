// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (= 30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°))
var zmin, zmax = math.NaN(), math.NaN()

func svg(w io.Writer) {
	zmin, zmax = minmax()

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if infCheck(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(az, bz, cz, dz), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func main() {
	svg(os.Stdout)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := computeSurface(i, j)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}
	return
}

func color(numbers ...float64) string {
	min := min(numbers...)
	max := max(numbers...)
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		return fmt.Sprintf("#%02x0000", int(red))
	}
	blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
	if blue > 255 {
		blue = 255
	}
	return fmt.Sprintf("#0000%02x", int(blue))
}

func min(numbers ...float64) float64 {
	min := math.NaN()
	for _, num := range numbers {
		if math.IsNaN(min) || num < min {
			min = num
		}
	}
	return min
}

func max(numbers ...float64) float64 {
	max := math.NaN()
	for _, num := range numbers {
		if math.IsNaN(max) || num > max {
			max = num
		}
	}
	return max
}

func corner(i, j int) (float64, float64, float64) {
	x, y, z := computeSurface(i, j)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func computeSurface(i, j int) (x float64, y float64, z float64) {
	// Find point (x, y) at corner of cell (i, j).
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)

	return x, y, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func infCheck(vals ...float64) bool {
	var isInf bool
	for _, val := range vals {
		isInf = isInf || math.IsInf(val, -1) || math.IsInf(val, +1)
	}
	return isInf
}
