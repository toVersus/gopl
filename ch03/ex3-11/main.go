package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	var out bytes.Buffer
	begin := 0
	if s[0] == '+' || s[0] == '-' {
		out.WriteByte(s[0])
		begin = 1
	}
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}

	target := s[begin:end]
	pre := len(target) % 3
	if pre > 0 {
		out.Write([]byte(target[:pre]))
		if len(target) > pre {
			out.WriteString(",")
		}
	}
	for i, c := range target[pre:] {
		if i%3 == 0 && i != 0 {
			out.WriteString(",")
		}
		out.WriteRune(c)
	}
	return out.String()
}
