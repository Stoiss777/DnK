/**
 * Exercise 3.8:
 *     Rendering fractals at high zoom levels demands great arithmetic precision.
 *     Implement the same fractal using four different representations of numbers: complex64,
 *     complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package.
 *     Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
 *     rational numbers.) How do they compare in performance and memory usage? At what zoom
 *     levels do rendering artifacts become visible?
 * 
 *     The complex64 version. Maximum zoom in this version is 10⁻⁴. 
 *     At higher zoom, the precision is not enough to take points on the graph.
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
 * 
 *     TODO: Return to this task after the chapter 11
 */
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	// Interesting coordinates for the selected zoom (left top point)
	xmin, ymax = 0.3385463985376018, 0.573318830897212
	// Zoom
	zoom = 0.00025
	// zoom = 0.000025  // This zoom gives "pixel art"
	// Iteration count
	iterations = 1000
	// Image size result
	width, height = 1024, 1024
	// Image contrast
	contrast = 5
)


func main() {
	var xmax, ymin float32 = xmin + zoom, ymax - zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex64) color.Color {
	var v complex64
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}