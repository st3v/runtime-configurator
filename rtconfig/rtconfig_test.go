package rtconfig

import (
	"reflect"
	"strings"
	"testing"
)

var (
	thisRuntimeConfig = RuntimeConfig{
		Releases: theseReleases,
		Addons:   theseAddons,
		Tags:     theseTags,
	}

	thatRuntimeConfig = RuntimeConfig{
		Releases: thoseReleases,
		Addons:   thoseAddons,
		Tags:     thoseTags,
	}
)

func TestRuntimeConfigAdd(t *testing.T) {
	have := thisRuntimeConfig.Add(thatRuntimeConfig)

	want := RuntimeConfig{
		Releases: thisRuntimeConfig.Releases.union(thatRuntimeConfig.Releases),
		Addons:   thisRuntimeConfig.Addons.union(thatRuntimeConfig.Addons),
		Tags:     thisRuntimeConfig.Tags.union(thatRuntimeConfig.Tags),
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}
}

func TestRuntimeConfigRemove(t *testing.T) {
	have := thisRuntimeConfig.Remove(thatRuntimeConfig)

	want := RuntimeConfig{
		Releases: thisRuntimeConfig.Releases.substract(thatRuntimeConfig.Releases),
		Addons:   thisRuntimeConfig.Addons.substract(thatRuntimeConfig.Addons),
		Tags:     thisRuntimeConfig.Tags.substract(thatRuntimeConfig.Tags),
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected result, want: %+v, have: %+v", want, have)
	}
}

func TestRead(t *testing.T) {
	data := `
---
releases:
- name: some-release
  version: some-version
addons:
- name: some-addon
  jobs:
  - name: some-job
    release: some-release
  properties:
    foo: bar
  include:
    deployments:
    - include-deployment
    jobs:
    - name: include-job
      release: include-release
    stemcell:
    - os: include-os
tags:
  bar: baz
`

	have, err := Read(strings.NewReader(data))
	if err != nil {
		t.Fatalf("unexpected error reading runtime config: %v", err)
	}

	want := RuntimeConfig{
		Releases: releases{{Name: "some-release", Version: "some-version"}},
		Addons: addons{{
			Name: "some-addon",
			Jobs: jobs{{
				"name":    "some-job",
				"release": "some-release",
			}},
			Properties: map[string]interface{}{
				"foo": "bar",
			},
			Include: map[string]interface{}{
				"deployments": []interface{}{
					"include-deployment",
				},
				"jobs": []interface{}{map[interface{}]interface{}{
					"name":    "include-job",
					"release": "include-release",
				}},
				"stemcell": []interface{}{map[interface{}]interface{}{
					"os": "include-os",
				}},
			},
		}},
		Tags: tags{"bar": "baz"},
	}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("unexpected runtime config read, want: %+v, have: %+v", want, have)
	}
}
