package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("%d: %s\n", i, arg)
	}
	fmt.Print(s)
}
