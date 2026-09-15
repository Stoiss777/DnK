/**
 * Exercise 4.07:
 *     Modify reverse to reverse the characters of a
 *     []byte slice that represents a UTF-8-encoded string, in place.
 *     Can you do it without allocating new memory?
 *     
 *     
 *     The answer: it is possible to implement it without allocation new memory,
 *     but every time the swap runes have different amount of bytes
 *     we must shift bytes between them.
 */
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sources := [...]string {
		"Hello world!",
		"Привет, странник",
		"Привет, 世界!",
		"悪魔城伝説",
	}

	for i, source := range sources {
		result := string(reverse([]byte(source)))

		fmt.Printf("___ Sample %d ___\n", i+1)
		fmt.Printf("Original: %s\n", source)
		fmt.Printf("Reversed: %s\n", result)
	}
}

func reverse(s []byte) []byte {
	for i,j := 0,len(s)-1; i<j; {
		rightRune, rightRuneLen := utf8.DecodeRune(s[j:])
		if (rightRune == utf8.RuneError && rightRuneLen == 1) {
			j--
			continue
		}

		leftRune, leftRuneLen := utf8.DecodeRune(s[i:])
		shift := rightRuneLen-leftRuneLen
		if shift > 0 {
			copy(s[i+shift:], s[i:j])
		} else if shift < 0 {
			copy(s[i:], s[i-shift:j+rightRuneLen])

		}
		utf8.EncodeRune(s[i:], rightRune)
		utf8.EncodeRune(s[j+shift:], leftRune)
		i += leftRuneLen + shift
		j += shift - 1
	}

	return s
}
