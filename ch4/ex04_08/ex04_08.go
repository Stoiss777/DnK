package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	catOther = 0
	catControl = 1
	catDigit = 2
	catLetter = 3
	catMark = 4
	catNumber = 5
	catPunct = 6
	catSpace = 7
	catTotal = 8
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen[utf8.UTFMax + 1]int  // count of lengths of UTF-8 encodings
	var utfcat[catTotal]int         // count of lenghts of UTF-8 categories
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()  // return rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		utfcat[category(r)]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Print("\ncategory\tcount\n")
	fmt.Printf("control\t\t%d\n", utfcat[catControl])
	fmt.Printf("digit\t\t%d\n", utfcat[catDigit])
	fmt.Printf("letter\t\t%d\n", utfcat[catLetter])
	fmt.Printf("mark\t\t%d\n", utfcat[catMark])
	fmt.Printf("number\t\t%d\n", utfcat[catNumber])
	fmt.Printf("punct\t\t%d\n", utfcat[catPunct])
	fmt.Printf("space\t\t%d\n", utfcat[catSpace])
	fmt.Printf("other\t\t%d\n", utfcat[catOther])
}

func category(r rune) int {
	switch {
	case unicode.IsControl(r):
		return catControl
	case unicode.IsDigit(r):
		return catDigit
	case unicode.IsLetter(r):
		return catLetter
	case unicode.IsMark(r):
		return catMark
	case unicode.IsNumber(r):
		return catNumber
	case unicode.IsPunct(r):
		return catPunct
	case unicode.IsSpace(r):
		return catSpace
	default:
		return catOther
	}
}
