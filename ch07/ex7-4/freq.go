// FindLinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type HTMLReader struct {
	s string
}

func (r *HTMLReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	n = copy(b, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewHTMLReader(s string) *HTMLReader {
	return &HTMLReader{s}
}

// visit counts number of each tag found in the links and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	var contents []string
	for _, urlStr := range os.Args[1:] {
		res, err := http.Get(urlStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get contents from %s: %s", urlStr, err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get contents from %s: %s", urlStr, err)
			os.Exit(1)
		}
		res.Body.Close()

		contents = append(contents, string(body))
	}

	for _, content := range contents {
		doc, err := html.Parse(NewHTMLReader(content))
		if err != nil {
			fmt.Fprintf(os.Stderr, "html parse failed: %s", err)
			os.Exit(1)
		}

		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}
}
