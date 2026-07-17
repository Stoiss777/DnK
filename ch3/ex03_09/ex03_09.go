/**
 * Exercise 3.9:
 *     Write a web server that renders fractals and writes the image data to the client.
 *     Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.
 * 
 *     I used the float64 version as a basis. The float32 version doesn't have enough precision.
 *     The big.Float version is too slow.
 * 
 *     Zoom to iteration ratio table
 *     ------------------
 *     zoom  | iterations
 *     ------------------
 *     3     | 200-500
 *     0.1   | 500–1000
 *     10⁻³  | 1000–2000
 *     10⁻⁶  | 2000–5000
 *     10⁻¹² | 5000-10000
 *     10⁻¹⁴ | 10000+
 */
package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"log"
	"strconv"
)

const (
	// Default x and y (left top point)
	xdef, ydef = -2, +2
	// Default zoom
	zdef = 4
	// Image size result
	width, height = 1024, 1024
	// Image contrast
	contrast = 5
)


func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64);
		y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64);
		z, _ := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64);
		if (x <= 0) {
			x = xdef
		}
		if (y <= 0) {
			y = ydef
		}
		if (z <= 0) {
			z = zdef
		}
		display(w, x, y, z)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func display(out io.Writer, xmin, ymax, zoom float64) {
	var xmax, ymin float64 = xmin + zoom, ymax - zoom
	var iterations uint = 10000

	// Number of iterations depends of required zoom
	if zoom < 10e-12 {
		iterations = 10000
	} else if zoom < 10e-12 {
		iterations = 5000
	} else if zoom < 10e-6 {
		iterations = 2000
	} else if zoom < 10e-1 {
		iterations = 1000
	} else if zoom < 1 {
		iterations = 500
	} else {
		iterations = 200
	}

	// fmt.Printf("xmin=%f, ymax=%f, zoom=%f, i=%d\n", xmin, ymax, zoom, iterations)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// z := complex(x, y)
			img.Set(px, py, mandelbrot(x, y, iterations))
		}
	}

	png.Encode(out, img)
}

/**
 * Cause this is a web server, I decided to use a faster algorithm,
 * without complex numbers.
 */
func mandelbrot(px, py float64, iterations uint) color.Color {
	var x, y float64

	for n := uint(0); n < iterations; n++ {
		x2 := x * x
		y2 := y * y

		if x2 + y2 > 4 {
			return color.Gray{255 - uint8(contrast*n)}
		}

		y = 2*x*y + py
		x = x2 - y2 + px
	}

	return color.Black
}
