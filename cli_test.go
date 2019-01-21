package main

import (
	"reflect"
	"testing"
)

func toStrPointer(s string) *string {
	return &s
}

func TestConditionParser(t *testing.T) {
	conditions := []string{
		"name:foo",
		"name like:foo%",
		"name >:bar",
		"name >=:bar",
		"name <:baz",
		"name <=:baz",
		"nameEmpty:",
		"nameNull:null",
		"nameNotNull:null ",
		"nameKeySpace  :foo",
		"nameValueSpace: foo",
		"nameColon::",
	}

	truth := map[string]*string{
		"name":           toStrPointer("foo"),
		"name like":      toStrPointer("foo%"),
		"name >":         toStrPointer("bar"),
		"name >=":        toStrPointer("bar"),
		"name <":         toStrPointer("baz"),
		"name <=":        toStrPointer("baz"),
		"nameEmpty":      toStrPointer(""),
		"nameNull":       nil,
		"nameNotNull":    toStrPointer("null "),
		"nameKeySpace":   toStrPointer("foo"),
		"nameValueSpace": toStrPointer(" foo"),
		"nameColon":      toStrPointer(":"),
	}

	r := conditionParser(conditions)

	if !reflect.DeepEqual(r, truth) {
		t.Fatalf("failed test expected: %v\n got: %v", truth, r)
	}
}
