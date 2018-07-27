package main

import (
	"fmt"
	"log"
	"os"

	"github.com/toversus/gopl/ch05/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			worklist = append(worklist, f(item)...)
		}
	}
}

func crawl(urlStr string) []string {
	fmt.Println(urlStr)
	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
