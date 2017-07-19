package slops

import (
	"reflect"
	"testing"
)

type diffUseCase struct {
	ladd      bool
	left      string
	radd      bool
	right     string
	leftOnly  []string
	rightOnly []string
}

var diffTestScript = []diffUseCase{
	{false, "", false, "", []string{}, []string{}},
	{false, "", true, "", []string{}, []string{""}},
	{true, "", false, "", []string{}, []string{}},
	{true, "", true, "a", []string{""}, []string{"a"}},
	{true, "a", false, "", []string{""}, []string{}},
	{true, "b", true, "b", []string{""}, []string{}},
	{false, "", true, "c", []string{""}, []string{"c"}},
	{false, "", true, "d", []string{""}, []string{"c", "d"}},
}

func TestDiff(t *testing.T) {
	left := make([]string, 0)
	right := make([]string, 0)

	for uci, uc := range diffTestScript {

		if uc.ladd {
			left = append(left, uc.left)
		}

		if uc.radd {
			right = append(right, uc.right)
		}

		lo, ro := Diff(left, right)

		if !reflect.DeepEqual(lo, uc.leftOnly) {
			t.Error("At index", uci, "left excess", lo, "does not match expected", uc.leftOnly)
		}

		if !reflect.DeepEqual(ro, uc.rightOnly) {
			t.Error("At index", uci, "right excess", ro, "does not match expected", uc.rightOnly)
		}
	}
}
