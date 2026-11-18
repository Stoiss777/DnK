/**
 * Exercise 5.03:
 *     Write a function to print the contents of all text nodes in an HTML
 *     document tree. Do not descend into <script> or <style> elements,
 *     since their contents are not visible in a web browser.
 */
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex05_03: %v\n", err)
		os.Exit(1)
	}

	walk("", doc)
}

func walk(tag string, n *html.Node) {
	if n.Type == html.ElementNode {
		tag = n.Data
	} else if n.Type == html.TextNode {
		if (tag != "script" && tag != "noscript" && tag != "style") {
			text := strings.TrimSpace(n.Data)
			// Do not show nodes with spaces only
			if text != "" {
				fmt.Printf("<%s>%s</%s>\n\n", tag, strings.TrimSpace(n.Data), tag)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(tag, c)
	}
}