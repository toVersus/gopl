// ex3.12 determines if strings are anagrams of each other.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Fprintln(os.Stderr, "please specify two strings")
		os.Exit(1)
	}
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(a, b string) bool {
	amap := make(map[rune]int)
	for _, c := range a {
		amap[c]++
	}
	bmap := make(map[rune]int)
	for _, c := range b {
		bmap[c]++
	}

	for _, c := range b {
		if v, ok := amap[c]; !ok || v != bmap[c] {
			return false
		}
	}
	for _, c := range a {
		if v, ok := bmap[c]; !ok || v != amap[c] {
			return false
		}
	}
	return true
}
