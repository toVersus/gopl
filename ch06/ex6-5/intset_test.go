package intset

import (
	"reflect"
	"testing"
)

type pair struct {
	x, y IntSet
}

func (p *pair) assert(t *testing.T, want []int) {
	t.Helper()

	for _, val := range want {
		if !p.x.Has(val) {
			t.Errorf("%d not found in %v", val, &p.x)
		}
	}
}

func TestHas(t *testing.T) {
	env := initTestEnv()
	want := []int{1, 9, 144}

	env.assert(t, want)
}

func TestString(t *testing.T) {
	env := initTestEnv()
	want := `{1 9 144}`

	result := env.x.String()
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got=%v, want=%v", result, want)
	}
}

func TestUnionWith(t *testing.T) {
	env := initTestEnv()
	want := []int{1, 9, 144}

	env.x.UnionWith(&env.y)

	env.assert(t, want)
}

func TestIntersectWith(t *testing.T) {
	env := initTestEnv()
	want := []int{1, 144}

	env.x.IntersectWith(&env.y)

	env.assert(t, want)
}

func TestDifferenceWith(t *testing.T) {
	env := initTestEnv()
	want := []int{9}

	env.x.DifferenceWith(&env.y)

	env.assert(t, want)
}

func TestSymmetricDifference(t *testing.T) {
	env := initTestEnv()
	want := []int{9}

	env.x.DifferenceWith(&env.y)

	env.assert(t, want)
}

func TestRemove(t *testing.T) {
	env := initTestEnv()
	want := []int{9}

	env.x.Remove(1)
	env.x.Remove(144)

	env.assert(t, want)
}

func TestElems(t *testing.T) {
	env := initTestEnv()
	want := []int{1, 9, 144}

	result := env.x.Elems()
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got=%v, want=%v", result, want)
	}
}

func initTestEnv() *pair {
	var env pair
	env.x.Add(1)
	env.x.Add(9)
	env.x.Add(144)

	env.y.Add(1)
	env.y.Add(144)

	return &env
}
