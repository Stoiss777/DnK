/**
 * Exercise 3.8:
 *     Rendering fractals at high zoom levels demands great arithmetic precision.
 *     Implement the same fractal using four different representations of numbers: complex64,
 *     complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package.
 *     Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
 *     rational numbers.) How do they compare in performance and memory usage? At what zoom
 *     levels do rendering artifacts become visible?
 * 
 *     The complex128 version. Maximum zoom in this version is 10⁻¹². 
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
	xmin, ymax = 0.33856375930859617, 0.5733096840391279
	// xmin, ymax = 0.338563759308993, 0.573309684039089  // interesting point for 1e-14
	// Zoom
	zoom = 5.58e-12
	// zoom = 3.58e-14  // This zoom gives "pixel art"
	// Iteration count
	iterations = 5000
	// Image size result
	width, height = 1024, 1024
	// Image contrast
	contrast = 5
)


func main() {
	var xmax, ymin float64 = xmin + zoom, ymax - zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}
