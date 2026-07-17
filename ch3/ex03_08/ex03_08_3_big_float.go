/**
 * Exercise 3.8:
 *     Rendering fractals at high zoom levels demands great arithmetic precision.
 *     Implement the same fractal using four different representations of numbers: complex64,
 *     complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package.
 *     Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
 *     rational numbers.) How do they compare in performance and memory usage? At what zoom
 *     levels do rendering artifacts become visible?
 * 
 *     The big.Float version. Zoom can by anything, but the larger the zoom,
 *     the longer the program execution time. In general, calculations for 
 *     the big.Float type take longer than calculations for the float type.
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
	"math/big"
	"os"
)

const (
	// Interesting coordinates for the selected zoom (left top point)
	xmin, ymax = 0.338563759308993, 0.573309684039089
	// Zoom
	zoom = 3.58e-14
	// Iteration count
	iterations = 10000
	// Image size result
	width, height = 1024, 1024
	// Float precision for big.Float
	precision = 128
	// Image contrast
	contrast = 5
)

// To speed up calculations, I made the initialization of these variables global.
var xx = new(big.Float).SetPrec(precision)
var yy = new(big.Float).SetPrec(precision)
var sum = new(big.Float).SetPrec(precision)
var tmp = new(big.Float).SetPrec(precision)
var two = new(big.Float).SetPrec(precision).SetFloat64(2)
var four = new(big.Float).SetPrec(precision).SetFloat64(4)


func main() {
	bigWidth := new(big.Float).SetPrec(precision).SetFloat64(width)
	bigHeight := new(big.Float).SetPrec(precision).SetFloat64(height)

	bigZoom := new(big.Float).SetPrec(precision).SetFloat64(zoom)
	bigXmin := new(big.Float).SetPrec(precision).SetFloat64(xmin)
	bigYmax := new(big.Float).SetPrec(precision).SetFloat64(ymax)
	bigXmax := new(big.Float).SetPrec(precision).Add(bigXmin, bigZoom)
	bigYmin := new(big.Float).SetPrec(precision).Sub(bigYmax, bigZoom)

	x := new(big.Float).SetPrec(precision)
	y := new(big.Float).SetPrec(precision)

	bigPx := new(big.Float).SetPrec(precision)
	bigPy := new(big.Float).SetPrec(precision)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		// Math here: y = float64(py)/height*(ymax-ymin) + ymin
		bigPy.SetFloat64(float64(py))
		y.Sub(bigYmax, bigYmin)
		y.Mul(bigPy, y)
		y.Quo(y, bigHeight)
		y.Add(y, bigYmin)

		for px := 0; px < width; px++ {
			// Math here: x = float64(px)/width*(xmax-xmin) + xmin
			bigPx.SetFloat64(float64(px))
			x.Sub(bigXmax, bigXmin)
			x.Mul(bigPx, x)
			x.Quo(x, bigWidth)
			x.Add(x, bigXmin)

			img.Set(px, py, mandelbrot(x, y))
		}
	}

	png.Encode(os.Stdout, img)
}

/**
 * - Since we are not using complex numbers here, I used a different version
 * of the algorithm.
 * 
 * - We haven't studied pointers yet, but I need to pass "big.Float" type 
 * parameters here.
 */
func mandelbrot(px, py *big.Float) color.Color {
	x := new(big.Float).SetPrec(precision)
	y := new(big.Float).SetPrec(precision)

	for n := 0; n < iterations; n++ {
        xx.Mul(x, x)
        yy.Mul(y, y)

		//  Math here: xx + yy > 4
		sum.Add(xx, yy)
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
		y.Copy(tmp)
	}

	return color.Black
}
