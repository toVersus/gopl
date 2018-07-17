package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}},
	}
	for _, test := range tests {
		if rotate(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("got=%q, want=%q\n", test.input, test.want)
		}
	}
}
