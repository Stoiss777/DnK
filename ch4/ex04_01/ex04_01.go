/**
 * Exercise 4.01:
 *     Write a function that counts the number of bits that are different in two SHA256
 *     hashes. (See PopCount from Section 2.6.2.)
 */
package main

import
(
	"crypto/sha256"
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	// For tests
	// var c1, c2 [32]byte
	// c1[0] = 2
	// c2[0] = 6

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Printf("Different bits: %d\n", diff(c1, c2))
}

func diff(c1, c2 [32]uint8) int {
	var cnt int

	for i:=0; i<32; i++ {
		cnt += popcount.PopCount(uint64(c1[i] ^ c2[i]))
	}

	return cnt
}