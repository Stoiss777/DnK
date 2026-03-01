/**
 * Exercise 1.3: 
 *     Experiment to measure the difference in running time between our
 *     potentially inefficient versions and the one that uses strings.Join.
 *     (Section 1.6 illustrates part of the time package,
 *     and Section 11.4 shows how to write benchmark tests
 *     for systematic performance evaluation.)
 */

// TODO: Finish it later!
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string

	start := time.Now()

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}