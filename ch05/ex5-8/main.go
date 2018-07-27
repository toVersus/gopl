package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func ElementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				fmt.Println("ok")
				return false
			}
		}
		return true
	}
	return forEachNode(n, pre, nil)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
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
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: URL ElementID")
		os.Exit(1)
	}
	urlStr := os.Args[1]
	id := os.Args[2]

	r := strings.NewReader(urlStr)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTML parser: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", ElementByID(doc, id))
}
