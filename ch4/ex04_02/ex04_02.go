/**
 * Exercise 4.02:
 *     Write a program that prints the SHA256 hash of its standard input by default but
 *     supports a command-line flag to print the SHA384 or SHA512 hash instead.
 */
package main

import
(
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

const (
	FlagSHA256 = iota
	FlagSHA384
	FlagSHA512
)

func main() {
	var val string
	var flag uint8 = FlagSHA256

	for i, arg := range os.Args[1:] {
		if i == 0 && arg[0:2] == "--" {
			switch (strings.ToLower(arg[2:])) {
			case "sha384":
				flag = FlagSHA384
			case "sha512":
				flag = FlagSHA512
			}
		} else {
			val = arg
		}
	}

	fmt.Printf("Input string: %s\n", val)

	switch flag {
	case FlagSHA384:
		c := sha512.Sum384([]byte(val))
		fmt.Printf("SHA384 hash: %x\nType: %T\n", c, c)
	case FlagSHA512:
		c := sha512.Sum384([]byte(val))
		fmt.Printf("SHA512 hash: %x\nType: %T\n", c, c)
	default:
		c := sha256.Sum256([]byte(val))
		fmt.Printf("SHA256 hash: %x\nType: %T\n", c, c)
	}
}
