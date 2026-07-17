/**
 * Exercise 3.10:
 *     Write a non-recursive version of comma, using bytes.Buffer instead of string
 *     concatenation.
 */
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("12345678"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		j := n - i - 1
		if j > 0 && j%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
