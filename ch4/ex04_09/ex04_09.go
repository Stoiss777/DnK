/**
 * Exercise 4.09:
 *     Write a program wordfreq to report the frequency of each word
 *     in an input text file. Call input.Split(bufio.ScanWords) before
 *     the first call to Scan to break the input into words instead of lines.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Print("Count\tWord\n-----\t-----\n")
	for word, count := range counts {
		fmt.Printf("%d\t%s\n", count, word)
	}
}
