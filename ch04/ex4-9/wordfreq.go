// ex4-9 counts and reports the frequency of each word in an input text file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

var filepath = flag.String("f", "input.txt", "specify a filepath")

type Word struct {
	Val  string
	Freq int
}

func main() {
	flag.Parse()

	f, err := os.Open(*filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n %s\n", *filepath, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	dict := make(map[string]int)
	for sc.Scan() {
		dict[sc.Text()]++
	}

	words := make([]Word, 0, len(dict))
	for val, freq := range dict {
		words = append(words, Word{val, freq})
	}
	sort.SliceStable(words, func(i, j int) bool {
		return words[i].Freq > words[j].Freq
	})

	for _, word := range words {
		fmt.Printf("%s:\t%d\n", word.Val, word.Freq)
	}
}
