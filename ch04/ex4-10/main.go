package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/toversus/gopl/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	format := "#%-5d %9.9s %.55s\n"

	dayBefore := make([]*github.Issue, 0)
	monthBefore := make([]*github.Issue, 0)
	yearBefore := make([]*github.Issue, 0)

	now := time.Now()
	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)
	for _, item := range result.Items {

		if item.CreatedAt.After(day) {
			dayBefore = append(dayBefore, item)
		}
		if item.CreatedAt.After(month) {
			monthBefore = append(monthBefore, item)
		}
		if item.CreatedAt.After(year) {
			yearBefore = append(yearBefore, item)
		}
	}

	fmt.Println("")
	fmt.Println("Past day:")
	for _, item := range dayBefore {
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}

	fmt.Println("")
	fmt.Println("Past month:")
	for _, item := range monthBefore {
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}

	fmt.Println("")
	fmt.Println("Past year:")
	for _, item := range yearBefore {
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}
