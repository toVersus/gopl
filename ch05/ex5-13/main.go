package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

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

var origHost string

func save(urlStr string) error {
	url, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid url: %s", err)
	}
	if len(origHost) == 0 {
		origHost = url.Host
	}
	if origHost != url.Host {
		return nil
	}

	var dir, filename string
	if len(filepath.Ext(filename)) == 0 {
		dir = filepath.Join(origHost, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(origHost, filepath.Dir(url.Path))
		filename = url.Path
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	if _, err = io.Copy(file, resp.Body); err != nil {
		return err
	}
	file.Close()

	return nil
}

func crawl(urlStr string) []string {
	fmt.Println(urlStr)
	if err := save(urlStr); err != nil {
		log.Printf("can't cache %s: %s", urlStr, err)
	}
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
