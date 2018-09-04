// Echo3 prints its command-line argument in more efficiently way.
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(echo(os.Args[1:]))
}
