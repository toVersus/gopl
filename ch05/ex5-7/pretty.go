package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int
var out bytes.Buffer

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.CommentNode:
		startComment(n)
	case html.TextNode:
		startText(n)
	}
}

func startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s=%s`, a.Key, a.Val))
	}
	var attrStr string
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := n.Data

	out.WriteString(fmt.Sprintf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end))
	depth++

}

func startComment(n *html.Node) {
	out.WriteString(fmt.Sprintf("%*s<!-- %s -->\n", depth*2, "", n.Data))
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	out.WriteString(fmt.Sprintf("%*s%s\n", depth*2, "", n.Data))
}

func end(n *html.Node) {
	if n.Type == html.ElementNode {
		endElement(n)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		out.WriteString(fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data))
	}
}

func main() {
	out = bytes.Buffer{}
	for _, urlStr := range os.Args[1:] {
		r := strings.NewReader(urlStr)
		doc, err := html.Parse(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTML parser: %v\n", err)
			continue
		}
		forEachNode(doc, start, end)
	}
	fmt.Println(out.String())
}
