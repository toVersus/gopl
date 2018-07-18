package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"A": 0}
	y := map[string]int{"B": 42}
	fmt.Println(equal(x, y))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
