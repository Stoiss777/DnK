/**
 * Exercise 3.3:
 *     Color each polygon based on its height, so that the peaks are colored
 *     red (#ff0000) and the valleys blue (#0000ff).
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
	angle         = math.Pi / 6 		// angle of x, y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)  // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height)

    // We have to make two passes to create the figure.
    // The first pass is to compute the possble values of zmin and zmax.
    zmin, zmax := zrange()

    for i := 0; i < cells; i++ {
    	for j := 0; j < cells; j++ {
    		ax, ay, az := corner(i+1, j)
    		bx, by, bz := corner(i, j)
    		cx, cy, cz := corner(i, j+1)
    		dx, dy, dz := corner(i+1, j+1)
    		// calculate the avarage z and find out the color
    		red, green, blue := color((az+bz+cz+dz)/4, zmin, zmax)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='none'" +
            	" stroke='#%02x%02x%02x'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, red, green, blue)
    	}
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r) / r
}

func zrange() (float64, float64) {
	zmin := math.Inf(1)
	zmax := math.Inf(-1)
    for i := 0; i < cells; i++ {
    	for j := 0; j < cells; j++ {
		    x := xyrange * (float64(i)/cells - 0.5)
		    y := xyrange * (float64(j)/cells - 0.5)
		    z := f(x, y)
			if z < zmin {
				zmin = z
			}
			if z > zmax {
				zmax = z
			}
    	}
    }
    return zmin, zmax
}

func color(z, zmin, zmax float64) (byte, byte, byte) {
	red := byte((z - zmin) / (zmax - zmin) * 0xFF)
	blue := 0xFF - red
	return red, 0, blue
}
