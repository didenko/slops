// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

import (
	"reflect"
	"testing"
)

type uniqueUseCase struct {
	in     []string
	expect []string
}

var uniqueTestScript = []uniqueUseCase{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{""}, []string{""}},
	{[]string{"-"}, []string{"-"}},
	{[]string{"-", "-"}, []string{"-"}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "c", "c"}, []string{"a", "b", "c"}},
}

func TestUnique(t *testing.T) {
	for uci, uc := range uniqueTestScript {
		result := Unique(uc.in)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}
