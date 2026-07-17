/**
 * Exercise 3.1:
 *     If the function f returns a non-finite float64 value,
 *     the SVG file will contain invalid <polygon>
 *      elements (although many SVG renderers handle this gracefully).
 *      Modify the program to skip invalid polygons.”
 */
package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 300            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)  // sin(30°), cos(30°)

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height)

    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    // x, y = 0, 0;  // Values for test the func f()
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    // The math.Hypot() method can return two types of non-finite values.
    // Zero is also dangerous to use, due to division by 'r' below.
    if r == 0 || math.IsNaN(r) || math.IsInf(r, 1) {
        return 0
    }
    return math.Sin(r) / r
}