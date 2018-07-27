package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$\w+|\${\w+}`)
var keyStr = flag.String("key", "foo", "specify a key which will be replaced by value")
var valStr = flag.String("val", "bar", "specify a value which will be used to replace key")

func expand(s string, f func(string) string) string {
	repl := func(s string) string {
		if strings.HasPrefix(s, "${") {
			s = s[2 : len(s)-1]
		} else {
			s = s[1:]
		}
		return f(s)
	}
	return pattern.ReplaceAllStringFunc(s, repl)
}

func main() {
	flag.Parse()

	subs := make(map[string]string, 0)
	k, v := *keyStr, *valStr
	subs[k] = v

	f := func(s string) string {
		v, ok := subs[s]
		if !ok {
			return "$" + s
		}
		return v
	}

	var out bytes.Buffer
	out.ReadFrom(os.Stdin)
	fmt.Print(expand(out.String(), f))
}
