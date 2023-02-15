package goption

import (
	"reflect"
	"testing"
)

type myStruct struct {
	name string
	tags []string
	m    map[string]myStruct
}

func TestSome(t *testing.T) {
	some1 := Some(25)
	requireEquals(t, "Some", some1, Option[int]{value: pointer(25)})

	some2 := Some("an string")
	requireEquals(t, "Some", some2, Option[string]{value: pointer("an string")})

	some3 := Some(myStruct{
		name: "the name",
		tags: []string{"a list", "of", "tags"},
		m: map[string]myStruct{
			"nested": {name: "nested_type", tags: []string{"this", "is", "nested"}, m: nil},
		},
	})
	requireEquals(t, "Some", some3, Option[myStruct]{value: pointer(myStruct{
		name: "the name",
		tags: []string{"a list", "of", "tags"},
		m: map[string]myStruct{
			"nested": {name: "nested_type", tags: []string{"this", "is", "nested"}, m: nil},
		},
	})})
}

func TestNone(t *testing.T) {
	none := None[string]()
	requireEquals(t, "None", none, Option[string]{})
}

func TestCopy(t *testing.T) {
	some := Some(myStruct{
		name: "a name",
		tags: []string{"hello"},
		m: map[string]myStruct{
			"another": {
				name: "nested", tags: []string{"this is a nested", "tag"},
			},
		},
	})

	copied := some.Copy()
	copied.Mutate(func(t *myStruct) {
		t.m = nil
		t.tags = append(t.tags, "world")
	})

	requireNotEquals(t, "Copy", some, copied)
}

func requireEquals(t *testing.T, name string, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}

func requireNotEquals(t *testing.T, name string, a, b interface{}) {
	if reflect.DeepEqual(a, b) {
		t.Errorf("%s = %v, dont want %v", name, a, b)
	}
}

func pointer[T any](value T) *T {
	return &value
}
