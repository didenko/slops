// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

import (
	"reflect"
	"testing"
)

type commonUseCase struct {
	left   []string
	right  []string
	expect []string
}

var commonTestScript = []commonUseCase{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"", "-"}, []string{}},
	{[]string{""}, []string{""}, []string{""}},
	{[]string{""}, []string{"", ""}, []string{""}},
	{[]string{"", ""}, []string{"", ""}, []string{"", ""}},
	{[]string{"a"}, []string{"b"}, []string{}},
	{[]string{"a", "b", "c"}, []string{"b", "d"}, []string{"b"}},
	{[]string{"a", "b", "c"}, []string{"a", "b"}, []string{"a", "b"}},
	{[]string{"b", "c"}, []string{"a", "b", "c"}, []string{"b", "c"}},
	{[]string{"a", "b", "c"}, []string{"a", "c"}, []string{"a", "c"}},
}

func TestCommon(t *testing.T) {
	for uci, uc := range commonTestScript {
		result := Common(uc.left, uc.right)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}
