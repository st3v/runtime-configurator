package rtconfig

import (
	"reflect"
	"testing"
)

var (
	theseAddons = addons{
		{Name: "foo", Jobs: jobs{{"foo": "1"}}},
		{Name: "bar", Jobs: jobs{{"bar": "1"}}},
	}

	thoseAddons = addons{
		{Name: "bar", Jobs: jobs{{"new": "2"}}},
		{Name: "baz", Jobs: jobs{{"baz": "2"}}},
	}
)

func TestAddonUnion(t *testing.T) {
	have := theseAddons.union(thoseAddons)

	want := addons{
		{Name: "foo", Jobs: jobs{{"foo": "1"}}},
		{Name: "bar", Jobs: jobs{{"new": "2"}}},
		{Name: "baz", Jobs: jobs{{"baz": "2"}}},
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseAddons) {
		t.Errorf("original value changed: %+v", theseAddons)
	}
}

func TestAddonSubstract(t *testing.T) {
	have := theseAddons.substract(thoseAddons)

	want := addons([]addon{
		{Name: "foo", Jobs: jobs{{"foo": "1"}}},
	})

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseAddons) {
		t.Errorf("original value changed: %+v", theseAddons)
	}
}
