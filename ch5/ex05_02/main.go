/**
 * Exercise 5.02:
 *     Write a function to populate a mapping from element names—p, div,
 *     span, and so on—to the number of elements with that name
 *     in an HTML document tree.
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex05_02: %v\n", err)
		os.Exit(1)
	}

	tags := make(map[string]int)
	calc(tags, doc)
	for tag, c := range tags {
		fmt.Printf("%s = %d\n", tag, c)
	}
}

func calc(tags map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		tags[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		calc(tags, c)
	}	
}