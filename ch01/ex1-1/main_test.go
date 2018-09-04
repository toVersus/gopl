package main

import (
	"reflect"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{
			args: []string{"main.go", "Hello,", "World!"},
			want: "main.go Hello, World!",
		},
	}

	for _, test := range tests {
		if result := echo(test.args); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
