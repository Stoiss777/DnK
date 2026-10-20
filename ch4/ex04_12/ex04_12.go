/**
 * Exercise 4.12:
 *     The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
 *     produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and
 *     build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript
 *     of each comic that matches a search term provided on the command line. 
 * 
 * 
 * The program is divided into three files for ease of reading.
 * 
 * Examples of running the program:
 *   xkcd download all  # re-downloads all JSON files
 *   xkcd download 571  # re-download only the JSON for comic 571
 *   xkcd index         # create a new search index
 *   xkcd               # interactive search mode
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Comix struct {
	Month string `json:"month"`
	Num int `json:"num"`
	Link string `json:"link"`
	Year string `json:"year"`
	News string `json:"news"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt string `json:"alt"`
	Img string `json:"img"`
	Title string `json:"title"`
	Day string `json:"day"`
}


func main() {
	var index Index
	var err error

	switch {
	case len(os.Args) > 2 && os.Args[1] == "download":
		if os.Args[2] == "all" {
			c, err := removeAll()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(2)				
			}
			fmt.Printf("%d file(s) deleted.\n", c)
			c, err = downloadAll(true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(2)				
			}
			fmt.Printf("%d file(s) downloaded.\n", c)
		} else {
			num, err := strconv.Atoi(os.Args[2])
			if err != nil || num < 1 {
				printSyntax()
				os.Exit(1)
			}
			exist, err := downloadComix(num, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(2)
			}
			if !exist {
				fmt.Fprintf(os.Stderr, "Not found!\n")
				os.Exit(2)
			}
		}

	case len(os.Args) > 1 && os.Args[1] == "index":
		index, err := buildIndex()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(3)
		}
		fmt.Printf("Found %d tokens.\n", len(index))
		saveIndex(IndexFile, index)

	case len(os.Args) == 1:
		index, err = loadIndex(IndexFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(3)
		}
		if index == nil {
			index, err = buildIndex()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
			saveIndex(IndexFile, index)
			fmt.Printf("Builded new index. Found %d tokens.\n", len(index))
		}	

		fmt.Print("Query to search: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		query := input.Text()

		ids := search(query, index)
		if len(ids) == 0 {
			fmt.Println("Nothing found.")
		}

		for _, id := range ids {
			comix, err := loadComix(fmt.Sprintf(ComixFile, id))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				break
			}
			fmt.Println()
			fmt.Printf("TITLE: %s\n", comix.Title)
			fmt.Printf("URL: " + ComixURL + "\n", comix.Num)
			fmt.Println("----------------------------")
			fmt.Println(comix.Transcript)
			fmt.Println()
		}
	default:
		printSyntax()
		os.Exit(1)
	}
}


func printSyntax() {
	fmt.Println("usage: xkcd [download|index] [number|all]")
}