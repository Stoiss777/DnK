/**
 * Exercise 3.5:
 *     Implement a full-color Mandelbrot set using the function image.NewRGBA and
 *     the type color.RGBA or color.YCbCr.
 */
package main

import (
	"image"
	"image/color"
	"image/png"
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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOT: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	
	// Iteration levels at which the image looks most interesting
	const pick1, pick2 = 20, 40

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n >= pick2 {
				return color.RGBA{255 - contrast*n, 0, 0, 0xff}
			} else if n >= pick1 {
				return color.RGBA{0, 255 - contrast*n, 0, 0xff}
			} else {
				return color.RGBA{0, 0, 255 - contrast*n, 0xff}
			}
		}
	}
	return color.Black
}
