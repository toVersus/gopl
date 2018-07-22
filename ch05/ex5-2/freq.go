// FindLinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// visit counts number of each tag found in the links and returns the result.
func visit(r io.Reader) map[string]int {
	freq := make(map[string]int, 0)
	z := html.NewTokenizer(os.Stdin)
	for type_ := z.Next(); type_ != html.ErrorToken; type_ = z.Next() {
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}

	return freq
}

func main() {
	for tag, count := range visit(os.Stdin) {
		fmt.Printf("%4d %s\n", count, tag)
	}
}
