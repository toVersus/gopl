package main

import (
	"fmt"
	"os"
	"strings"
)

func echo(args []string) string {
	var s string
	for i, arg := range args {
		s += fmt.Sprintf("%d: %s\n", i, arg)
	}
	return strings.TrimRight(s, "\n")
}

func main() {
	fmt.Println(echo(os.Args[1:]))
}
