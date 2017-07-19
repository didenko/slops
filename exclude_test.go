package slops

import (
	"reflect"
	"testing"
)

type excludeUseCase struct {
	source  []string
	rejects []string
	expect  []string
}

var excludeTestScript = []excludeUseCase{
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
}

func TestExclude(t *testing.T) {
	for uci, uc := range excludeTestScript {
		result := Exclude(uc.source, uc.rejects)

		if !reflect.DeepEqual(uc.expect, result) {
			t.Error("At index", uci, "result", result, "does not match expected", uc.expect)
		}
	}
}
