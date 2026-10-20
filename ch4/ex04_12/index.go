/**
 * This file contains functions that implement an inverted index.
 * 
 */
package main

import (
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

type Index map[string]map[int]struct{}

const IndexFile = "index.dat"


func search(query string, index Index) []int {
	var result []int

	tokens := tokenize(query)

	for id := range index[tokens[0]] {
		found := true

		for _, token := range tokens[1:] {
			if _, ok := index[token][id]; !ok {
				found = false
				break
			}
		}

		if found {
			result = append(result, id)
		}
	}

	slices.Sort(result)

	return result
}


func buildIndex() (Index, error) {
	files, err := filepath.Glob(ComicGlob)
	if err != nil {
		return nil, err
	}

	index := make(Index)

	for _, file := range files {
		comic, err := loadComic(file)
		if err != nil {
			return nil, err
		}
		for _, word := range tokenize(comic.Title + " " + comic.Transcript) {
			if index[word] == nil {
				index[word] = make(map[int]struct{})
			}
			index[word][comic.Num] = struct{}{}
		}
	}

	return index, nil
}


func saveIndex(filename string, index Index) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	return gob.NewEncoder(f).Encode(index)
}


func loadIndex(filename string) (Index, error) {
    f, err := os.Open(filename)
    if errors.Is(err, os.ErrNotExist) {
    	return nil, nil
    } else if err != nil {
    	return nil, err
    }

    var index Index

    if err := gob.NewDecoder(f).Decode(&index); err != nil {
        return nil, err
    }

    return index, nil
}


func tokenize(s string) []string {
	return strings.FieldsFunc(
		strings.ToLower(s),
		func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		},
	)
}
