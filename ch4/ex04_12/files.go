/**
 * This file contains functions that implement the loading, saving, and reading of JSON files for comics.
 * 
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const ComicURL = "https://xkcd.com/%d/"
const ComicJSON = "https://xkcd.com/%d/info.0.json"
const ComicFile = "storage/%d.json"
const ComicGlob = "storage/*.json"
const ComicMaxNumber = 9999


func loadComic(filename string) (*Comic, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var encoded Comic
	if err := json.Unmarshal(data, &encoded); err != nil {
		return nil, err
	}

	return &encoded, nil
}


func downloadComic(number int, verbose bool) (bool, error) {
	if verbose {
		fmt.Printf("Comic %d downloading... ", number)
	}
	resp, err := http.Get(fmt.Sprintf(ComicJSON, number))
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return false, err
	}

	if err = os.WriteFile(fmt.Sprintf(ComicFile, number), data, 0666); err != nil {
		resp.Body.Close()
		return false, err
	}
	if verbose {
		fmt.Printf("OK.\n")
	}

	return true, nil
}


func downloadAll(verbose bool) (int, error) {
	var err error
	var exist bool
	var c int = 0
	for i:=1; i<ComicMaxNumber; i++ {
		// xkcd joke
		if i == 404 {
			continue
		}
		exist, err = downloadComic(i, verbose)
		if err != nil {
			return c, err
		}
		if !exist {
			if verbose {
				fmt.Printf("No more items found.\n")
			}
			return c, nil
		}
		c++
	}
	return c, nil
}


func removeAll() (int, error) {
    files, err := filepath.Glob(ComicGlob)
    if err != nil {
        return 0, err
    }

    var c int = 0
    for _, file := range files {
        if err := os.Remove(file); err != nil {
            return c, err
        }
        c++
    }

    return c, nil
}
