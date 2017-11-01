// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

import (
	"reflect"
	"testing"
)

type excludeUseCase struct {
	source  []string
	rejects []string
	expect  []string
}

var excludeAllTestScript = []excludeUseCase{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"", "-"}, []string{}},
	{[]string{"0"}, []string{}, []string{"0"}},
	{[]string{"0"}, []string{"0"}, []string{}},
	{[]string{"1"}, []string{"0"}, []string{"1"}},
	{[]string{"a", "b"}, []string{"a", "b"}, []string{}},
	{[]string{"a", "b"}, []string{"a", "a"}, []string{"b"}},
	{[]string{"a", "a"}, []string{"a"}, []string{}},
	{[]string{"a", "b", "b", "c"}, []string{"b", "c"}, []string{"a"}},
	{[]string{"a", "b"}, []string{"b", "c"}, []string{"a"}},
	{[]string{"b", "c"}, []string{"a", "b"}, []string{"c"}},
	{[]string{"b", "c", "c"}, []string{"a", "b"}, []string{"c", "c"}},
}

func TestExcludeAll(t *testing.T) {
	for uci, uc := range excludeAllTestScript {
		result := ExcludeAll(uc.source, uc.rejects)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}

func TestExcludeAllContrived(t *testing.T) {
	for uci, uc := range excludeAllTestScript {
		result := ExcludeAllContrived(uc.source, uc.rejects)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}

var excludeByCountTestScript = []excludeUseCase{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"", "-"}, []string{}},
	{[]string{"0"}, []string{}, []string{"0"}},
	{[]string{"0"}, []string{"0"}, []string{}},
	{[]string{"1"}, []string{"0"}, []string{"1"}},
	{[]string{"a", "b"}, []string{"a", "b"}, []string{}},
	{[]string{"a", "b"}, []string{"a", "a"}, []string{"b"}},
	{[]string{"a", "a"}, []string{"a"}, []string{"a"}},                     // differs from ExcludeAll
	{[]string{"a", "b", "b", "c"}, []string{"b", "c"}, []string{"a", "b"}}, // differs from ExcludeAll
	{[]string{"a", "b"}, []string{"b", "c"}, []string{"a"}},
	{[]string{"b", "c"}, []string{"a", "b"}, []string{"c"}},
	{[]string{"b", "c", "c"}, []string{"a", "b"}, []string{"c", "c"}},
}

func TestExcludeByCount(t *testing.T) {
	for uci, uc := range excludeByCountTestScript {
		result := CollectExcludedByCount(uc.source, uc.rejects, getAll)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}

var benchSrc = []string{
	"./slops",
	"./slops",
	"./slops",
	"./slops",
	"./slops",
	"./slops/.travis.yml",
	"./slops/collector.go",
	"./slops/common.go",
	"./slops/common_test.go",
	"./slops/diff.go",
	"./slops/diff_test.go",
	"./slops/exclude.go",
	"./slops/exclude_test.go",
	"./slops/exclude_test.go",
	"./slops/exclude_test.go",
	"./slops/exclude_test.go",
	"./slops/exclude_test.go",
	"./slops/merge.go",
	"./slops/merge_test.go",
	"./slops/merge_test.go",
	"./slops/unique.go",
	"./slops/unique_test.go",
	"./slops/variety.go",
	"./slops/variety_test.go",
}

var benchRejects = []string{
	"./slops",
	"./slops",
	"./slops/.travis.yml",
	"./slops/exclude_test.go",
	"./slops/merge_test.go",
	"./slops/merge_test.go",
	"./slops/unknown_test.go",
}

// Needed so that compiler optimizations do not eliminate benchmarks
var benchResult []string

func BenchmarkExcludeAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchResult = ExcludeAll(benchSrc, benchRejects)
	}
}

func BenchmarkExcludeAllContrived(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchResult = ExcludeAllContrived(benchSrc, benchRejects)
	}
}
