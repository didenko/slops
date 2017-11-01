// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

import (
	"reflect"
	"testing"
)

type diffUseCase struct {
	left      []string
	right     []string
	onlyLeft  []string
	onlyRight []string
}

var diffTestScript = []diffUseCase{
	{nil, nil, []string{}, []string{}},
	{[]string{}, nil, []string{}, []string{}},
	{[]string{}, []string{}, []string{}, []string{}},
	{[]string{}, []string{"-"}, []string{}, []string{"-"}},
	{[]string{"-"}, []string{"-"}, []string{}, []string{}},
	{[]string{"-", "-"}, []string{"-"}, []string{"-"}, []string{}},
	{[]string{"-", "-", "-"}, []string{"-"}, []string{"-", "-"}, []string{}},
	{[]string{"-", "-"}, []string{"-", "a"}, []string{"-"}, []string{"a"}},
	{[]string{"-", "-", "a"}, []string{"-", "a"}, []string{"-"}, []string{}},
	{[]string{"-", "-", "a", "b"}, []string{"-", "a", "b"}, []string{"-"}, []string{}},
	{[]string{"-", "-", "b"}, []string{"-", "a", "b"}, []string{"-"}, []string{"a"}},
	{[]string{"-", "b"}, []string{"-", "a", "b"}, []string{}, []string{"a"}},
	{[]string{"-", "b"}, []string{"-", "a", "b", "c"}, []string{}, []string{"a", "c"}},
}

func TestDiff(t *testing.T) {
	for uci, uc := range diffTestScript {

		lo, ro := Diff(uc.left, uc.right)

		if !reflect.DeepEqual(lo, uc.onlyLeft) {
			t.Error("At index", uci, "left excess", lo, "does not match expected", uc.onlyLeft)
		}

		if !reflect.DeepEqual(ro, uc.onlyRight) {
			t.Error("At index", uci, "right excess", ro, "does not match expected", uc.onlyRight)
		}
	}
}
