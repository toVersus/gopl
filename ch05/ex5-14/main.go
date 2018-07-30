package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"linear algebra": {"calculus"},
	"calculus":       {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal language":       {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

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

func main() {
	// get random course
	var course string
	for course = range prereqs {
		break
	}
	deps := func(course string) []string {
		fmt.Println(course)
		return prereqs[course]
	}
	breadthFirst(deps, []string{course})
}
