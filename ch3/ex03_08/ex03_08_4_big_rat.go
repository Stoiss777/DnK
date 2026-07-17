/**
 * Exercise 3.8:
 *     Rendering fractals at high zoom levels demands great arithmetic precision.
 *     Implement the same fractal using four different representations of numbers: complex64,
 *     complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package.
 *     Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
 *     rational numbers.) How do they compare in performance and memory usage? At what zoom
 *     levels do rendering artifacts become visible?
 * 
 *     The bit.Rat version. Zoom can by anything, but the calculation takes a huge amount of time.
 *     After just a small number of iterations, the fraction denominators begin to grow very quickly
 *     and can contain thousands or tens of thousands of bits, and each next operation becomes increasingly
 *     more expensive. The Apple M1 processor was unable to calculate a 1024x1024 pixel graph at 4x scaling
 *     even over 24 hours. To make the calculations take at least some acceptable time, I added rounding of 
 *     the fraction at each iteration in the function "mandelbrot".
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
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

const (
	// Iteration count
	iterations = 100
	// Precision (for rounding)
	precision = 8
	// Image contrast
	contrast = 5
)

// Interesting coordinates for the selected zoom (left top point)
var xmin = big.NewRat(-2, 1)
var ymax = big.NewRat(2, 1)
// Because the calculations take a huge amount of time,
// I use a large zoom in this sample.
var zoom = big.NewRat(4, 1)
// Image size result
var width = big.NewRat(1024, 1)
var height = big.NewRat(1024, 1)

// To speed up calculations, I made the initialization of these variables global.
var two = big.NewRat(2, 1)
var four = big.NewRat(4, 1)
var sum = new(big.Rat)
var tmp = new(big.Rat)
var xx = new(big.Rat)
var yy = new(big.Rat)


func main() {
	w := int(width.Num().Uint64())
	h := int(height.Num().Uint64())

	xmax := new(big.Rat).Add(xmin, zoom)
	ymin := new(big.Rat).Sub(ymax, zoom)

	x := new(big.Rat)
	y := new(big.Rat)

	bx := new(big.Rat)
	by := new(big.Rat)

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for py := 0; py < h; py++ {
		// Math here: y = float64(py)/height*(ymax-ymin) + ymin
		by.SetInt64(int64(py))
		y.Sub(ymax, ymin)
		y.Mul(by, y)
		y.Quo(y, height)
		y.Add(y, ymin)

		for px := 0; px < w; px++ {
			// Math here: x = float64(px)/width*(xmax-xmin) + xmin
			bx.SetInt64(int64(px))
			x.Sub(xmax, xmin)
			x.Mul(bx, x)
			x.Quo(x, width)
			x.Add(x, xmin)

			// fmt.Printf("px=%d py=%d x=%s y=%s\n", px, py, x.RatString(), y.RatString())

			img.Set(px, py, mandelbrot(x, y))
		}
	}

	png.Encode(os.Stdout, img)
}

/**
 * - Since we are not using complex numbers here, I used a different version
 * of the algorithm.
 * 
 * - We haven't studied pointers yet, but I need to pass "big.Rat" type 
 * parameters here.
 */
func mandelbrot(px, py *big.Rat) color.Color {
	x := big.NewRat(0, 1)
	y := big.NewRat(0, 1)

	for n := 0; n < iterations; n++ {
        xx.Mul(x, x)
        yy.Mul(y, y)

        // This code rounds the numerator and denominator.
        // This reduces accuracy but speeds up calculations.
        // If you need unlimited precision, you need to delete these two lines,
        // but then the calculations will take a huge amount of time.
		xx.SetString(xx.FloatString(precision))
		yy.SetString(yy.FloatString(precision))

		//  Math here: xx + yy > 4
		sum.Add(xx, yy)

		// fmt.Printf("\tn=%d xx=%s yy=%s\n", n, xx.RatString(), yy.RatString())
		if sum.Cmp(four) > 0 {
			return color.Gray{255 - uint8(contrast*n)}
		}

		//  Math here: tmp = 2*x*y + py
		tmp.Mul(x, y)
		tmp.Mul(tmp, two)
		tmp.Add(tmp, py)
		//  Math here: x = xx - yy + px
		x.Sub(xx, yy)
		x.Add(x, px)
		// y = t
		y.Set(tmp)
	}

	return color.Black
}
