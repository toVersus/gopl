package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"one", "two", "three", "two", "one"}, []string{"one", "two", "three"}},
		{[]string{"a", "b", "b", "c", "b", "a", "c"}, []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		if result := unique(test.input); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q, want=%q\n", result, test.want)
		}
	}
}
