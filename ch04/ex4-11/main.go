package main

import (
	"fmt"
	"log"
	"os"

	"github.com/toversus/gopl/ch04/ex4-11/github"
)

var usage string = `usage
search QUERY
[get|close|open] OWNER REPO ISSUE_NUMBER
update OWNER REPO ISSUE_NUMBER TITLE BODY
`

func main() {
	if len(os.Args) < 2 {
		log.Fatalln(usage)
	}

	action := os.Args[1]
	args := os.Args[2:]
	if action == "search" {
		if len(args) < 1 {
			log.Fatalln(usage)
		}
		github.SearchIssues(args)
		return
	}

	if len(args) != 3 {
		log.Fatalln(usage)
	}
	owner, repo, number := args[0], args[1], args[2]
	switch action {
	case "get":
		issue, err := github.GetIssue(owner, repo, number)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
			owner, repo, number, issue.User.Login, issue.Title, issue.Body)
	case "update":
		if len(args) != 5 {
			log.Fatalln(usage)
		}
		title, body := args[3], args[4]
		github.UpdateIssue(owner, repo, number, title, body)
	case "close":
		github.CloseIssue(owner, repo, number)
	case "open":
		github.OpenIssue(owner, repo, number)
	}
}
