// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

import (
	"reflect"
	"testing"
)

type mergeUseCase struct {
	left   []string
	right  []string
	expect []string
}

var mergeTestScript = []mergeUseCase{
	{[]string{}, []string{}, []string{}},

	{[]string{"a"}, []string{}, []string{"a"}},
	{[]string{"a"}, []string{"b"}, []string{"a", "b"}},
	{[]string{"a", "b"}, []string{"b", "c"}, []string{"a", "b", "c"}},

	{[]string{}, []string{"a"}, []string{"a"}},
	{[]string{"b"}, []string{"a"}, []string{"a", "b"}},
	{[]string{"b", "c"}, []string{"a", "b"}, []string{"a", "b", "c"}},

	{[]string{"a", "a", "a"}, []string{"-", "a", "a"}, []string{"-", "a", "a", "a"}},
	{[]string{"-", "a", "a"}, []string{"a", "a", "a"}, []string{"-", "a", "a", "a"}},
}

func TestMerge(t *testing.T) {
	for uci, uc := range mergeTestScript {
		result := Merge(uc.left, uc.right)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}

var mergeUniqueTestScript = []mergeUseCase{
	{[]string{}, []string{}, []string{}},

	{[]string{"a"}, []string{}, []string{"a"}},
	{[]string{"a"}, []string{"b"}, []string{"a", "b"}},
	{[]string{"a", "b"}, []string{"b", "c"}, []string{"a", "b", "c"}},

	{[]string{}, []string{"a"}, []string{"a"}},
	{[]string{"b"}, []string{"a"}, []string{"a", "b"}},
	{[]string{"b", "c"}, []string{"a", "b"}, []string{"a", "b", "c"}},

	{[]string{"a", "a", "a"}, []string{"-", "a", "a"}, []string{"-", "a"}},
	{[]string{"-", "a", "a"}, []string{"a", "a", "a"}, []string{"-", "a"}},
}

func TestMergeUnique(t *testing.T) {
	for uci, uc := range mergeTestScript {
		result := Merge(uc.left, uc.right)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}
