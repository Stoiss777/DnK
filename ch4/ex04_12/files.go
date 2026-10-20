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

const ComixURL = "https://xkcd.com/%d/"
const ComixJSON = "https://xkcd.com/%d/info.0.json"
const ComixFile = "storage/%d.json"
const ComixGlob = "storage/*.json"
const ComixMaxNumber = 9999


func loadComix(filename string) (*Comix, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var encoded Comix
	if err := json.Unmarshal(data, &encoded); err != nil {
		return nil, err
	}

	return &encoded, nil
}


func downloadComix(number int, verbose bool) (bool, error) {
	if verbose {
		fmt.Printf("Comix %d downloading... ", number)
	}
	resp, err := http.Get(fmt.Sprintf(ComixJSON, number))
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

	if err = os.WriteFile(fmt.Sprintf(ComixFile, number), data, 0666); err != nil {
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
	for i:=1; i<ComixMaxNumber; i++ {
		// xkcd joke
		if i == 404 {
			continue
		}
		exist, err = downloadComix(i, verbose)
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
    files, err := filepath.Glob(ComixGlob)
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
