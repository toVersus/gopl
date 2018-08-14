package counter

import (
	"bytes"
	"reflect"
	"testing"
)

type counterTestCase struct {
	input []byte
	want  int64
}

func TestCountingWriter(t *testing.T) {
	tests := []counterTestCase{
		{[]byte("Hello, world!"), 13},
		{[]byte(""), 0},
	}

	for _, test := range tests {
		r := bytes.NewBuffer([]byte(""))
		b := &byteCounter{r, 0}
		c, n := CountingWriter(b)
		c.Write(test.input)
		if !reflect.DeepEqual(*n, test.want) {
			t.Errorf("got=%d, want=%d", *n, test.want)
		}
	}
}
