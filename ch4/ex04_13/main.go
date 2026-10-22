/**
 * Exercise 4.13:
 *     The JSON-based web service of the Open Movie Database
 *     lets you search https://omdbapi.com/ for a movie by name and
 *     download its poster image. Write a tool poster that downloads
 *     the poster image for the movie named on the command line.
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

const ApiUrl = "https://omdbapi.com/demo.aspx/?t=%s&token=demo"

type Movie struct {
	Title string
	Year string
	Genre string
	Director string
	Poster string
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use the second paramter as a search query")
		os.Exit(1)
	}

	movie, err := search(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf(
		"Title: %s\nYear: %s\nGenre: %s\nDirector: %s\n",
		movie.Title, movie.Year, movie.Genre, movie.Director,
	)

    if movie.Poster == "N/A" || movie.Poster == "" {
		fmt.Fprintf(os.Stderr, "Poster not available.\n")
		os.Exit(0)
    }
    u, err := url.Parse(movie.Poster)
    if err != nil {
		fmt.Fprintf(os.Stderr, "Bad poster URL: %s.\n", movie.Poster)
		os.Exit(3)
    }
    ext := path.Ext(u.Path)
    if ext == "" {
		fmt.Fprintf(os.Stderr, "Unknown poster format: %s.\n", movie.Poster)
		os.Exit(3)
    }
    filename := movie.Title + ext

	fmt.Print("Downloading poster... ")
	err = download(movie.Poster, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(4)
	}
	fmt.Printf("saved to \"%s\"\n", filename)
}


func search(query string) (*Movie, error) {
	resp, err := http.Get(
		fmt.Sprintf(ApiUrl, url.QueryEscape(query)),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	movie := new(Movie)
	if err := json.NewDecoder(resp.Body).Decode(movie); err != nil {
		resp.Body.Close()
		return nil, err
	}
	if (movie.Title == "") {
		resp.Body.Close()
		return nil, fmt.Errorf("movie not found")
	}

	resp.Body.Close()
	return movie, nil
}


func download(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	file, err := os.Create(filename)
	if err != nil {
		resp.Body.Close()
		return err
	}

	_, err = io.Copy(file, resp.Body)

	resp.Body.Close()
	file.Close()
	return err
}
