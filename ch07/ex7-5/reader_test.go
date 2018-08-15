package reader

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	input string
	limit int64
	want  string
}

func TestLimitReader(t *testing.T) {
	tests := []testCase{
		{"Hello world!", 4, "Hell"},
	}

	for _, test := range tests {
		r := LimitReader(strings.NewReader(test.input), test.limit)
		result, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(string(result), test.want) {
			t.Errorf("got=%s, want=%s", string(result), test.want)
		}
	}
}
