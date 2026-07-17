/**
 * Exercise 3.11:
 *     Enhance comma so that it deals correctly with floating-point numbers and an
 *     optional sign.
 */
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("-123"))
	fmt.Println(comma("-123456.93889"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("error"))
	fmt.Println(comma("err12345.123"))
	fmt.Println(comma("12345.err123"))
}

func comma(s string) string {
	var buf bytes.Buffer
	var i, j int

	slen := len(s)

	// sign (in general, everything that comes before the numbers)
	for i=0; i<slen && s[i]<'0' && s[i]>'9'; i++ {
		buf.WriteByte(s[i])
	}

	// integer part
	for j = i; j<slen && s[j]>='0' && s[j]<='9'; j++ {}
	intbuff := s[i:j]
	bufflen := len(intbuff)
	// add commas to the integer part
	for i := 0; i < bufflen; i++ {
		buf.WriteByte(intbuff[i])
		j := bufflen - i - 1
		if j > 0 && j%3 == 0 {
			buf.WriteByte(',')
		}
	}

	// the remaining part of the number
	for ; j<slen; j++ {
		buf.WriteByte(s[j])
	}


	return buf.String()	
}
