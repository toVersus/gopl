package main

import "fmt"

func main() {
	var stack []int
	stack = append(stack, 5, 6, 7, 8, 9)
	s := stack
	fmt.Println(remove2(stack, 2))
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// remove2 doesn't preseve the order of elements.
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
