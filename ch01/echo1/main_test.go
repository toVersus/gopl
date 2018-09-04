package main

import (
	"reflect"
	"testing"
)

func TestEcho1(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{
			args: []string{"<source file>", "Hello,", "World!"},
			want: "Hello, World!",
		},
	}

	for _, test := range tests {
		if result := echo(test.args); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
