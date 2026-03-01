/**
 * Exercise 1.4: 
 *     Modify dup2 to print the names of all files in which
 *     each duplicated line occurs.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// index - a string from input,
	// value - list of affected files separated by "\n"
	// It would be more correct to use a two-dimensional array here,
	// but we haven't studied that yet at this point.
	// But we learnt about the strings.Split() function.
	faffected := make(map[string]string)
	// Final list of affected files
	fsummary := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, "", counts, faffected)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, arg, "dup2: %v\n", err)

				continue
			}
			countLines(f, arg, counts, faffected)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			// The string always starts with "\n",
			// so the first element is always an empty string and can be omitted.
			for _, fname := range strings.Split(faffected[line][1:], "\n") {
				if fname != "" {
					fsummary[fname]++
				}
			}
		}
	}

	if (len(fsummary) > 0) {
		fmt.Println("\nAffected files:")
		for fname, _ := range fsummary {
			fmt.Printf("%s\n", fname)
		}
	}
}

func countLines(f *os.File, fname string, counts map[string]int, faffected map[string]string) {
	var line string
	input := bufio.NewScanner(f)
	for input.Scan() {
		line = input.Text()
		counts[line]++
		faffected[line] += "\n" + fname
	}
	// NOTE: ignoring potential errors from input.Err()
}