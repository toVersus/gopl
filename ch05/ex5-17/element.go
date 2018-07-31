package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ElementByTagName(n *html.Node, tags ...string) []*html.Node {
	nodes := []*html.Node{}
	keep := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		if _, ok := keep[tag]; !ok {
			keep[tag] = struct{}{}
		}
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		if _, ok := keep[n.Data]; ok {
			nodes = append(nodes, n)
		}
		return true
	}
	forEachNode(n, pre, nil)
	return nodes
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := []*html.Node{}
	u = append(u, n)
	for len(u) > 0 {
		n, u = u[0], u[1:]
		if pre != nil && !pre(n) {
			return n
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}

		if post != nil && !post(n) {
			return n
		}
	}
	return nil
}

func main() {
	for _, urlStr := range os.Args[1:] {
		r := strings.NewReader(urlStr)
		doc, err := html.Parse(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "element: %v\n", err)
			os.Exit(1)
		}
		for _, n := range ElementByTagName(doc, "img") {
			fmt.Printf("%+v\n", n)
		}
	}
}
