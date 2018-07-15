package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Fprintln(os.Stderr, "please specify a integer")
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var out bytes.Buffer
	for i := 0; i < n; i += 3 {
		out.WriteString(",")
		out.WriteString(s[i : i+3])
	}
	return out.String()[1:]
}
