// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func echo(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func main() {
	fmt.Println(echo(os.Args))
}
