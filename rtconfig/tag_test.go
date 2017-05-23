package rtconfig

import (
	"reflect"
	"testing"
)

var (
	theseTags = tags{
		"foo": "1",
		"bar": "1",
	}

	thoseTags = tags{
		"bar": "2",
		"baz": "2",
	}
)

func TestTagsUnion(t *testing.T) {
	have := theseTags.union(thoseTags)

	want := tags{
		"foo": "1",
		"bar": "2",
		"baz": "2",
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseTags) {
		t.Errorf("original value changed: %+v", theseTags)
	}
}

func TestTagsSubstract(t *testing.T) {
	have := theseTags.substract(thoseTags)

	want := tags{"foo": "1"}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}

	if reflect.DeepEqual(have, theseTags) {
		t.Errorf("original value changed: %+v", theseTags)
	}
}
