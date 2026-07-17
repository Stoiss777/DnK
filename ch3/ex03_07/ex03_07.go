/**
 * Exercise 3.7:
 *     Another simple fractal uses Newton's method to find complex solutions to a
 *     function such as z^4-1=0. Shade each starting point by the number of iterations required to
 *     get close to one of the four roots. Color each point by the root it approaches.
 */
package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
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
            // img.Set(px, py, mandelbrot(z))
            img.Set(px, py, newton(z))
        }
    }
    png.Encode(os.Stdout, img) // NOT: ignoring errors
}

func newton(z complex128) color.Color {
    const (
        iterations = 100
        tolerance = 1e-6
    )
    // The four roots of z^4-1=0: 1, -1, i, -i (pre-calculated manually)
    var roots = [4]complex128{1, -1, 1i, -1i}
    // Color for each root
    var colors = [4][3]float64{  // ???
        {220, 50, 50},   // red   — root  1
        {50, 100, 220},  // blue  — root -1
        {50, 200, 80},   // green — root  i
        {220, 180, 50},  // gold  — root -i
    }

    for n := uint8(0); n < iterations; n++ {
        // Newton step: z = z - f(z)/f'(z)
        // f(z)  = z^4 - 1
        // f'(z) = 4z^3
        z3 := z*z*z
        z = z - (z3*z - 1) / (4*z3)

        for r, root := range roots {
            if cmplx.Abs(z - root) < tolerance {
                // Darken pixels that took more iterations (boundary regions stay dark).
                max := math.Log(float64(iterations + 1))
                contrast := 1.0 - 0.75*math.Log(float64(n+1))/max

                return color.RGBA {
                    R: uint8(colors[r][0] * contrast),
                    G: uint8(colors[r][1] * contrast),
                    B: uint8(colors[r][2] * contrast),
                    A: 0xff,
                }
            }
        }
    }

    return color.Black
}
