// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Diff takes sorted slices of strings and returns the non-
// overlapping entries. The result is still sorted.
func Diff(left, right []string) (leftOnly, rightOnly []string) {
	ga := &getAll{}
	return CollectDifferent(left, right, ga, ga)
}

// CollectDifferent applies related Collectors to every item which is
// in one slice (left or right), but not in the other (right or left)
// slice. Both input slices are expected to be sorted.
func CollectDifferent(left, right []string, lc, rc Collector) (leftColl, rightColl []string) {

	leftColl = make([]string, 0)
	rightColl = make([]string, 0)

	for l, r := 0, 0; l < len(left) || r < len(right); {

		if l == len(left) || left[l] > right[r] {
			rightColl = rc.Collect(rightColl, right[r])
			r++
			continue
		}

		if r == len(right) || left[l] < right[r] {
			leftColl = lc.Collect(leftColl, left[l])
			l++
			continue
		}

		l++
		r++
	}

	return leftColl, rightColl
}
