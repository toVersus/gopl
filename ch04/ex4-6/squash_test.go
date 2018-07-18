package squash

import (
	"reflect"
	"testing"
)

func TestSquash(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("世界\t世界\n"), []byte("世界 世界 ")},
		{[]byte("\t\v\f\r\n"), []byte("     ")},
	}
	for _, test := range tests {
		if result := squash(test.input); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%#v, want=%#v\n", result, test.want)
		}
	}
}
