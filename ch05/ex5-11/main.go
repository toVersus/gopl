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

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(course string, items []string)

	visitAll = func(course string, items []string) {
		for _, item := range items {
			for _, prerp := range m[item] {
				if course == prerp {
					order = append(order, item)
				}
			}

			if seen[item] {
				continue
			}
			seen[item] = true
			visitAll(item, m[item])
			order = append(order, item)
		}
	}

	for key := range m {
		visitAll(key, []string{key})
	}

	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
