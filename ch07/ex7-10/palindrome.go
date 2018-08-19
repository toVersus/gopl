package main

import (
	"fmt"
	"sort"
)

type Palindrom []byte

func (x Palindrom) Len() int           { return len(x) }
func (x Palindrom) Less(i, j int) bool { return x[i] < x[j] }
func (x Palindrom) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	j := s.Len()
	for i := 0; i < j; i++ {
		j--
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(Palindrom([]byte("abcba"))))
}
