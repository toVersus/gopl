package main

import (
	"strings"
	"testing"
)

var args = []string{"arg1", "arg2", "arg3", "arg4", "arg5", "arg6", "arg7", "arg8", "arg9", "arg10"}

func inefficientEcho() string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
	}
	return s
}

var result string

func BenchmarkInefficientEcho(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = inefficientEcho()
	}
	result = s
}

func BenchmarkEcho(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = strings.Join(args, " ")
	}
	result = s
}
