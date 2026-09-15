/**
 * Exercise 4.05:
 *     Write an in-place function to eliminate adjacent duplicates in a []string slice.
 */
package main

import "fmt"

func main() {
	var res []string

	var example1 = []string {
		"The 1st string",
		"The 2nd string",
		"The 2nd string",
		"The 3rd string",
		"The 3rd string",
		"The 3rd string",
		"The 4th string",
	}
	res = dup(example1)
	for _, s := range res {
		fmt.Println(s)
	}

	fmt.Println("---")
	var example2 = []string {
		"The 1st string",
	}
	res = dup(example2)
	for _, s := range res {
		fmt.Println(s)
	}

	fmt.Println("---")
	var example3 = []string {}
	res = dup(example3)
	for _, s := range res {
		fmt.Println(s)
	}
}

func dup(source []string) []string {
	if len(source) == 0 {
		return source
	}

	out := source[:1]
	prev := source[0]

	for i:=1; i<len(source); i++ {
		if source[i] != prev {
			out = append(out, source[i])
			prev = source[i]
		}
	}

	return out
}