package counter

import (
	"reflect"
	"testing"
)

type counterTestCase struct {
	input []byte
	want  int
}

func TestByteCounter(t *testing.T) {
	tests := []counterTestCase{
		{[]byte("Hello, world!"), 13},
		{[]byte(""), 0},
	}

	for _, test := range tests {
		var b ByteCounter
		b.Write(test.input)

		if !reflect.DeepEqual(b, ByteCounter(test.want)) {
			t.Errorf("got=%d, want=%d", b, test.want)
		}
	}
}

func TestWordCounter(t *testing.T) {
	tests := []counterTestCase{
		{[]byte("Hello, world!"), 2},
		{[]byte(""), 0},
		{[]byte("one two three four five"), 5},
	}

	for _, test := range tests {
		var w WordCounter
		w.Write(test.input)

		if !reflect.DeepEqual(w, WordCounter(test.want)) {
			t.Errorf("got=%d, want=%d", w, test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	tests := []counterTestCase{
		{[]byte("Hello, world!"), 1},
		{[]byte(""), 0},
		{[]byte("one\ntwo\nthree\nfour\nfive\n"), 5},
	}

	for _, test := range tests {
		var l LineCounter
		l.Write(test.input)

		if !reflect.DeepEqual(l, LineCounter(test.want)) {
			t.Errorf("got=%d, want=%d", l, test.want)
		}
	}
}
