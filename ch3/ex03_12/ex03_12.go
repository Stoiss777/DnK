/**
 * Exercise 3.12:
 *     Write a function that reports whether two strings are anagrams of each other,
 *     that is, they contain the same letters in a different order.
 */
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(anagram("test", "bob"))
	fmt.Println(anagram("test", "sony"))
	fmt.Println(anagram("listen", "silent"))
	fmt.Println(anagram("listin", "silent"))
	fmt.Println(anagram("Earth", "Heart"))
	fmt.Println(anagram("Апельсин", "Спаниель"))
}

func anagram(s1, s2 string) bool {
	var i, j int

	b1 := []rune(s1)
	b2 := []rune(s2)
	n1 := len(b1)
	n2 := len(b2)

	if (n1 != n2) {
		return false
	}

	for i = 0; i < n1; i++ {
		for j = 0; j < n2; j++ {
			if unicode.ToLower(b1[i]) == unicode.ToLower(b2[j]) {
				b2[j] = 0
				break
			}
		}
		if (j == n2) {
			return false
		}
	}

	return true
}