/**
 * Exercise 4.14:
 *     Create a web server that queries GitHub once and then allows
 *     navigation of the list of bug reports, milestones, and users.
 */
package main

import (
	"fmt"
	"os"
	// "path/filepath"
	"html/template"
	"net/http"
	"log"
	"strconv"

	"gopl.io/ch4/github"
)

const (
	BindAddress = "localhost:8000"
	IndexTemplateFile = "templates/index.html.tmpl"
	BugTemplateFile = "templates/bug.html.tmpl"
	MilestoneTemplateFile = "templates/milestone.html.tmpl"
	ContributorTemplateFile = "templates/contributor.html.tmpl"
)

type appData struct {
	Repo, Owner string
    Bugs []github.Issue
    Milestones []github.Milestone
    Contributors []github.User	
}

func main() {
	var data appData
	var err error

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: ex04_14 repo_owner repo_name.\n")
		os.Exit(1)
	}

	fmt.Println("Github login...")
	user, err := github.Login()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Login failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s.\nFetching data... ", user.Login)

	data.Owner = os.Args[1]
	data.Repo = os.Args[2]
	data.Bugs, err = github.GetBugs(data.Owner, data.Repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	data.Milestones, err = github.GetMilestones(data.Owner, data.Repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	data.Contributors, err = github.GetContributors(data.Owner, data.Repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("OK\n\n")


    http.HandleFunc(
    	"/",
    	func(writer http.ResponseWriter, request *http.Request) {
			tmpl := template.Must(template.ParseFiles("templates/index.html.tmpl"))
	        writer.Header().Set("Content-Type", "text/html")
	        err = tmpl.Execute(writer, data)
	        if err != nil {
	        	fmt.Fprint(writer, "Template not found")
	        }
    	},
    )

    http.HandleFunc(
    	"/bugs/{id}",
    	func(writer http.ResponseWriter, request *http.Request) {
		    id, err := strconv.Atoi(request.PathValue("id"))
		    if err != nil {
		        http.Error(writer, "Invalid ID", http.StatusBadRequest)
		        return
		    }
		    bug := getBugByNumber(data.Bugs, id)
		    if bug == nil {
		        http.Error(writer, "Not found", http.StatusNotFound)
		        return
		    }
			tmpl := template.Must(template.ParseFiles(BugTemplateFile))
	        writer.Header().Set("Content-Type", "text/html")
	        err = tmpl.Execute(writer, bug)
	        if err != nil {
	        	http.Error(writer, "Template not found", http.StatusInternalServerError)
	        }
	    },
    )

    http.HandleFunc(
    	"/milestones/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
		    id, err := strconv.Atoi(request.PathValue("id"))
		    if err != nil {
		        http.Error(writer, "Invalid ID", http.StatusBadRequest)
		        return
		    }
		    milestone := getMilestoneByNumber(data.Milestones, id)
		    if milestone == nil {
		        http.Error(writer, "Not found", http.StatusNotFound)
		        return
		    }
			tmpl := template.Must(template.ParseFiles(MilestoneTemplateFile))
	        writer.Header().Set("Content-Type", "text/html")
	        err = tmpl.Execute(writer, milestone)
	        if err != nil {
	        	http.Error(writer, "Template not found", http.StatusInternalServerError)
	        }
	    },
    )

    http.HandleFunc(
    	"/contributors/{id}",
		func(writer http.ResponseWriter, request *http.Request) {
		    id, err := strconv.Atoi(request.PathValue("id"))
		    if err != nil {
		        http.Error(writer, "Invalid ID", http.StatusBadRequest)
		        return
		    }
		    contributor := getContributorByID(data.Contributors, id)
		    if contributor == nil {
		        http.Error(writer, "Not found", http.StatusNotFound)
		        return
		    }
			tmpl := template.Must(template.ParseFiles(ContributorTemplateFile))
	        writer.Header().Set("Content-Type", "text/html")
	        err = tmpl.Execute(writer, contributor)
	        if err != nil {
	        	http.Error(writer, "Template not found", http.StatusInternalServerError)
	        }
	    },
    )    

	fmt.Printf("HTTP Server started on http://%s\n", BindAddress)
    log.Fatal(http.ListenAndServe(BindAddress, nil))
}


func getBugByNumber(bugs []github.Issue, number int) *github.Issue {
	for _, bug := range bugs {
		if bug.Number == number {
			return &bug
		}
	}

	return nil
}

func getMilestoneByNumber(milestones []github.Milestone, number int) *github.Milestone {
	for _, milestone := range milestones {
		if milestone.Number == number {
			return &milestone
		}
	}

	return nil
}

func getContributorByID(contributors []github.User, ID int) *github.User {
	for _, contributor := range contributors {
		if contributor.ID == ID {
			return &contributor
		}
	}

	return nil
}