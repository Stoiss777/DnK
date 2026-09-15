/**
 * Exercise 4.03:
 *     Rewrite reverse to use an array pointer instead of a slice.
 */
package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)  // pass the array by reference
	fmt.Println(a)
}

// The number of elements must be specified because a classic array is used.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


/*
 * Original reverse
 *
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
*/