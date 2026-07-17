/**
 * Exercise 3.4:
 *     Following the approach of the Lissajous example in Section 1.7,
 *     construct a web server that computes surfaces and writes SVG data to
 *     the client.
 */
package main

import (
    "fmt"
    "io"
    "math"
    "net/http"
    "log"
    "strconv"
)

const (
    cells               = 100                 // number of grid cells
    xyrange             = 30.0                // axis ranges (-xyrange..+xyrange)
    angle               = math.Pi / 6         // angle of x, y axes
    minwidth, minheight = 600, 300            // minimum image resolution
    maxwidth, maxheight = 4096, 2160          // maximum image resolution
    defcolor            = 0                   // default figure's color '#000000'
)

var sin30, cos30    = math.Sin(angle), math.Cos(angle)  // sin(30°), cos(30°)
var width, height int  // canvas size in pixels
var color int64        // default color #000000
var xyscale float64    // pixels per x or y unit
var zscale float64     // pixels per z unit


func main() {
    handler := func(writer http.ResponseWriter, request *http.Request) {
        w, _ := strconv.Atoi(request.URL.Query().Get("width"))
        h, _ := strconv.Atoi(request.URL.Query().Get("height"))
        c, _ := strconv.ParseInt(request.URL.Query().Get("color"), 16, 32)
        // Default values values
        width, height, color = minwidth, minheight, defcolor
        // Check and validate user's input 
        if w > minwidth && w <= maxwidth {
            width = w
        }
        if h > minheight && h <= maxheight {
            height = h
        }
        if c > 0 && c <= 16777215 {  // possible RGB color range
            color = c
        }
        xyscale = float64(width) / 2 / xyrange 
        zscale = float64(height) * 0.4

        writer.Header().Set("Content-Type", "image/svg+xml")
        surface(writer)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer) {
    fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>\n", width, height)

    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='none'" +
                " stroke='#%06x'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, color)
        }
    }
    fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := float64(width)/2 + (x-y)*cos30*xyscale
    sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r) / r
}