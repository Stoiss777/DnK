package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopl.io/ch2/popcount"
)

// Hint: 18446744073709551615 - the biggest uint64 number

func main() {
	var start time.Time
	var elapsed float64

	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		start = time.Now()
		sum := popcount.PopCount(x)
		elapsed += time.Since(start).Seconds()
		fmt.Println(sum)
	}

	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())
}