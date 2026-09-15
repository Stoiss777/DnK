/**
 * Exercise 4.04:
 *    Write a version of rotate that operates in a single pass.
 */
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, shift uint) {
	for i, j := 0, shift; j < uint(len(s)); i, j = i+1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
}
