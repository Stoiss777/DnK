// The app converts its numeric argument to:
//   Celsius <-> Fahrenheit
//   Foot    <-> Meter
//   Pound   <-> Kilogram
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lengthconv"
	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Printf("Temp:\t%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		t := lengthconv.Foot(v)
		m := lengthconv.Meter(v)
		fmt.Printf("Length:\t%s = %s, %s = %s\n",
			t, lengthconv.FToM(t), m, lengthconv.MToF(m))

		p := weightconv.Pound(v)
		k := weightconv.Kilogram(v)
		fmt.Printf("Weight:\t%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}