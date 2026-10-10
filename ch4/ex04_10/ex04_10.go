/**
 * Exercise 4.10:
 *     Modify issues to report the results in age categories,
 *     say less than a month old, less than a year old, and more
 *     than a year old.
 */
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	runTime := time.Now()
	monthAgo := runTime.AddDate(0, -1, 0)
	yearAgo := runTime.AddDate(-1, 0, 0)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s\t%s\t%.55s\n", 
			item.Number, item.User.Login, ago(item.CreatedAt, monthAgo, yearAgo), item.Title)
	}
}

func ago(createdAt, monthAgo, yearAgo time.Time) string {
	switch {
	case createdAt.Compare(monthAgo) > 0:
		return "less than a month old"
	case createdAt.Compare(yearAgo) > 0:
		return "less than a year old"
	default:
		return "more that a year old"
	}
}