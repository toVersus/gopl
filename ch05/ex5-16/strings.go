package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Join is the another implementation of strings.Join with variadic arguments.
func Join(sep string, strs ...string) string {
	var buf bytes.Buffer
	if len(strs) == 0 {
		return ""
	}
	for _, str := range strs {
		buf.WriteString(str + sep)
	}
	return strings.TrimRight(buf.String(), sep)
}

func main() {
	fmt.Println(Join("\n", "a", "b", "c"))
	fmt.Println(Join(""))
}
