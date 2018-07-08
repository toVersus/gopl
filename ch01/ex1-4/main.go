package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filemapper := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filemapper)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filemapper)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%+v\n", n, line, filemapper[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filemapper map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		if !contains(filemapper[line], f.Name()) {
			filemapper[line] = append(filemapper[line], f.Name())
		}
	}
}

func contains(filenames []string, key string) bool {
	for _, filename := range filenames {
		if filename == key {
			return true
		}
	}
	return false
}
