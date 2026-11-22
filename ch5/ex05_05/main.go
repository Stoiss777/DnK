/**
 * Exercise 5.05:
 *     Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
 */
package main

import (
    "bufio"
	"fmt"
    "strings"
    "net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        words, images, err := CountWordsAndImages(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ex05_05: %v\n", err)
            continue
        }
        fmt.Printf("%d words, %d images.\n", words, images)
    }
}


// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages("", doc)
    return
}


func countWordsAndImages(tag string, n *html.Node) (words, images int) {
    if n.Type == html.TextNode && tag != "script" && tag != "noscript" && tag != "style" {
        input := bufio.NewScanner(strings.NewReader(n.Data))
        input.Split(bufio.ScanWords)
        for input.Scan() {
            // fmt.Printf("> %s\n", input.Text())
            words++
        }
    } else if n.Type == html.ElementNode {
        tag = n.Data
        if tag == "img" {
            images++
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(tag, c)
        words += w
        images += i
    }

    return
}