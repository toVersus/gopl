package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func fetch(stack []string, n *html.Node) {
	s := strings.Trim(n.Data, "\n")
	s = strings.Trim(s, "\t")
	if n.Type == html.TextNode && s != "" {
		stack = append(stack, s)
		fmt.Printf("%#v\n", stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fetch(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fetch(nil, doc)
}
