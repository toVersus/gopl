package graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	graph = map[string]map[string]bool{}
	tests := []struct {
		from, to string
		want     bool
	}{
		{"a", "b", true},
	}
	for _, test := range tests {
		addEdge(test.from, test.to)
		if result := hasEdge(test.from, test.to); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%t, want=%t\n", result, test.want)
		}
	}
}
