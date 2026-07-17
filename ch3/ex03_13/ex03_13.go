/**
 * Exercise 3.13:
 *     Write const declarations for KB, MB, up through YB as compactly as you can.
 */
package main

import (
	"fmt"
	"math/big"
)

/*
In the book says:
	"The iota mechanism has its limits. For example, it's not possible to 
	generate the more familiar powers of 1000 (KB, MB, and so son) because
	there is no exponentiation operator."

Thus, the shortest way to write constants is to simply write the numbers directly.
*/
const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {
	fmt.Printf("1 KB = %.0f bytes\n", KB)
	fmt.Printf("1 MB = %.0f bytes\n", MB)

	// The float64 type does not have enough precision to store YB.
	// The Printf function cannot display the correct value.
	fmt.Printf("1 YB = %.0f bytes\n", YB)
}
