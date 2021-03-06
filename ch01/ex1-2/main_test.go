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
			args: []string{"Hello,", "World!"},
			want: "0: Hello,\n1: World!",
		},
	}

	for _, test := range tests {
		if result := echo(test.args); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
