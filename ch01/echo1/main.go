// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func echo(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func main() {
	fmt.Println(echo(os.Args))
}
