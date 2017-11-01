// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

import (
	"reflect"
	"testing"
)

type varietyUseCase struct {
	left   []string
	right  []string
	expect []string
}

var varietyTestScript = []varietyUseCase{
	{[]string{}, []string{}, []string{}},

	{[]string{"a"}, []string{}, []string{"lo:a"}},
	{[]string{"a"}, []string{"b"}, []string{"lo:a", "ro:b"}},
	{[]string{"a", "b"}, []string{"b", "c"}, []string{"lo:a", "cm:b", "ro:c"}},

	{[]string{}, []string{"a"}, []string{"ro:a"}},
	{[]string{"b"}, []string{"a"}, []string{"ro:a", "lo:b"}},
	{[]string{"b", "c"}, []string{"a", "b"}, []string{"ro:a", "cm:b", "lo:c"}},

	{[]string{"a", "a", "a"}, []string{"-", "a", "a"}, []string{"ro:-", "cm:a", "cm:a", "lo:a"}},
	{[]string{"-", "a", "a"}, []string{"a", "a", "a"}, []string{"lo:-", "cm:a", "cm:a", "ro:a"}},
}

func addLabel(label string) Collector {
	return func(pile []string, item string) []string {
		return append(pile, label+item)
	}
}

func TestVariety(t *testing.T) {
	for uci, uc := range varietyTestScript {
		result := CollectVariety(uc.left, uc.right, addLabel("lo:"), addLabel("cm:"), addLabel("ro:"))

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}
