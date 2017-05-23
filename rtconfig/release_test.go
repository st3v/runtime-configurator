package rtconfig

import (
	"reflect"
	"testing"
)

var (
	theseReleases = releases{
		{Name: "foo", Version: "1"},
		{Name: "bar", Version: "1"},
	}

	thoseReleases = releases{
		{Name: "bar", Version: "2"},
		{Name: "baz", Version: "2"},
	}
)

func TestReleasesUnion(t *testing.T) {
	have := theseReleases.union(thoseReleases)

	want := releases{
		{Name: "foo", Version: "1"},
		{Name: "bar", Version: "2"},
		{Name: "baz", Version: "2"},
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseReleases) {
		t.Errorf("original value changed: %+v", theseReleases)
	}
}

func TestReleasesSubstract(t *testing.T) {
	have := theseReleases.substract(thoseReleases)

	want := releases{{Name: "foo", Version: "1"}}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseReleases) {
		t.Errorf("original value changed: %+v", theseReleases)
	}
}
