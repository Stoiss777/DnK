/**
 * Exercise 4.06:
 *     Write an in-place function that squashes each run of adjacent Unicode
 *     spaces (see unicode.IsSpace) in a UTF-8-encoded []byte
 *     slice into a single ASCII space.
 */
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)


func main() {
	var sample string = "  Hello\u2003\u2002world !  "
	res := sdup([]byte(sample))

	fmt.Printf("Original string: \"%s\", %d byte(s)\n", sample, len(sample))
	fmt.Printf("Result string: \"%s\", %d byte(s)\n", string(res), len(string(res)))
}

func sdup(source []byte) []byte {
	var isSpace, isPrevSpace bool = false, false
	out := source[:0]

	for len(source) > 0 {
		r, size := utf8.DecodeRune(source)
		if unicode.IsSpace(r) {
			isSpace = true
			r = rune(' ')
		} else {
			isSpace = false
		}
		if !(isSpace && isPrevSpace) {
			out = utf8.AppendRune(out, r)
		}
		isPrevSpace = isSpace
		source = source[size:]
	}

	return out
}