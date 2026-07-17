/**
 * Exercise 3.6:
 *     Supersampling is a technique to reduce the effect of pixelation by computing the
 *     color value at several points within each pixel and taking the average. The simplest method is
 *     to divide each pixel into four "subpixels". Implement it.
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
        scale                  = 2
    )

    srcImage := image.NewRGBA(image.Rect(0, 0, width * scale, height * scale))
    srcWidth := width * scale

    srcHeight := height * scale

    for py := 0; py < srcHeight; py++ {
        y := float64(py)/float64(srcHeight)*(ymax-ymin) + ymin
        for px := 0; px < srcWidth; px++ {
            x := float64(px)/float64(srcWidth)*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            srcImage.Set(px, py, mandelbrot(z))
        }
    }

    
    // The simplest supersampling algorithm the "box filter",
    // compute one pixel from four "subpixels".
    dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

    for py := 0; py < height; py++ {
        for px := 0; px < width; px++ {

            var r, g, b, a int
            for sy := 0; sy < scale; sy++ {
                for sx := 0; sx < scale; sx++ {
                    c := srcImage.RGBAAt(px*scale + sx, py*scale + sy)
                    r += int(c.R)
                    g += int(c.G)
                    b += int(c.B)
                    a += int(c.A)
                }
            }

            n := scale * scale
            dstImage.SetRGBA(px, py, color.RGBA{
                R: uint8(r / n),
                G: uint8(g / n),
                B: uint8(b / n),
                A: uint8(a / n),
            })
        }
    }

    png.Encode(os.Stdout, dstImage)
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
