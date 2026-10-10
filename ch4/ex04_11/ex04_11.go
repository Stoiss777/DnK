/**
 * Exercise 4.11:
 *     Build a tool that lets users create, read, update, and delete GitHub
 *     issues from the command line, invoking their preferred text editor
 *     when substantial text input is required.
 * 
 * 
 * Exit statuses:
 *     0 - Success
 *     1 - Wrong syntax
 *     2 - Login error
 *     3 - Getting issue error
 *     4 - Creating issue error
 *     5 - Updating issue error
 *     6 - Deleting issue error
 *     7 - Issue number not specified.
 *     8 - Editor error
 *     9 - Wrong syntax in an editor
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"gopl.io/ch4/github"
)

func main() {
	if len(os.Args) < 3 {
		printSyntax()
		os.Exit(1)
	}

	user, err := github.Login()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Login failed: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Hello %s\n\n", user.Login)

	params := &github.IssueRequestParams {
		Owner: user.Login,
		Repo: os.Args[2],
	}
	if len(os.Args) == 4 {
		params.Number, err = strconv.Atoi(os.Args[3])
		if err != nil {
			printSyntax()
			os.Exit(1)
		}
	}

	switch os.Args[1] {
	case "get":
		if params.Number == 0 {
			issues, err := github.GetIssues(params)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(3)
			}
			for _, issue := range issues {
				printIssueHeader(issue)
			}			
		} else {
			issue, err := github.GetIssue(params)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(3)
			}
			printIssue(issue)
		}

	case "create":
		input, err := editData(nil)
		if err != nil {
			fmt.Errorf("%v\n", err)
			os.Exit(8)
		}
		parts := bytes.Split(input, []byte("\n\n"))
		if len(parts) < 2 {
			fmt.Printf("Use \\n\\n to split title and description.")
			os.Exit(9)
		}

		data := &github.IssueRequestData{
			Title: string(parts[0]),
			Body: string(parts[1]),
			Assignees: []string{ user.Login },
		}
		issue, err := github.CreateIssue(params, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(4)
		}
		fmt.Println("Created issue:")
		printIssue(issue)

	case "update":
		if params.Number == 0 {
			fmt.Printf("Issue number not specified.")
			os.Exit(7)
		}
		issue, err := github.GetIssue(params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(5)
		}
		issueText := []byte(issue.Title + "\n\n"+ issue.Body)
		input, err := editData(issueText)
		if err != nil {
			fmt.Errorf("%v\n", err)
			os.Exit(8)
		}
		parts := bytes.Split(input, []byte("\n\n"))
		if len(parts) < 2 {
			fmt.Printf("Use \\n\\n to split title and description.")
			os.Exit(9)
		}
		data := &github.IssueRequestData{
			Title: string(parts[0]),
			Body: string(parts[1]),
		}
		issue, err = github.UpdateIssue(params, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(5)
		}
		fmt.Println("Updated issue:")
		printIssue(issue)

	case "delete":
		if params.Number == 0 {
			fmt.Printf("Issue number not specified.")
			os.Exit(7)
		}
		err := github.DeleteIssue(params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(6)
		}
		fmt.Println("Issue has been deleted.")

	default:
		printSyntax()
		os.Exit(1)	
	}

}


func editData(content []byte) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, fmt.Errorf("EDITOR is not set")
	}

	tmp, err := os.CreateTemp("", "draft-*.txt")
	if err != nil {
		return nil, err
	}
	if content != nil {
		if _, err = tmp.Write(content); err != nil {
			return nil, err
		}
	}
	filename := tmp.Name()
	if err := tmp.Close(); err != nil {
		os.Remove(filename)
		return nil, err
	}

	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Remove(filename)
		return nil, err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		os.Remove(filename)
		return nil, err
	}

	return data, nil
}


func printIssue(issue *github.Issue) {
	printIssueHeader(issue)
	fmt.Println("-------")
	fmt.Println(issue.Body)
}

func printIssueHeader(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s\t%.55s\n", 
		issue.Number, issue.User.Login, issue.Title)
}

func printSyntax() {
	fmt.Println("usage: issue get|create|update|delete repo [number]")
}