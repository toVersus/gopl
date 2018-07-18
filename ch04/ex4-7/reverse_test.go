package reverse

import (
	"reflect"
	"testing"
)

func TestRevUTF8(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{[]byte("Hello, 世界"), "界世 ,olleH"},
	}
	for _, test := range tests {
		if result := string(revUTF8(test.input)); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got %v, want %v", result, test.want)
		}
	}
}
