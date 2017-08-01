// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Common gathers same entries from two sorted slices into
// a new slice. The order is preserved. The lesser number of
// duplicates is preserved
func Common(left, right []string) []string {
	return CollectCommon(left, right, getAll)
}

// CommonUnique gathers same entries from two sorted slices into
// a new slice. The order is preserved. Duplicates are reduced to
// a single item
func CommonUnique(left, right []string) []string {
	return CollectCommon(left, right, getUnique)
}

// CollectCommon applies a Collector to every item which is
// in both left and right slices. Both input slices are
// expected to be sorted.
func CollectCommon(left, right []string, collect Collector) []string {

	common := make([]string, 0)

	for i, j := 0, 0; i < len(left) && j < len(right); {
		if left[i] < right[j] {
			i++
			continue
		}

		if left[i] > right[j] {
			j++
			continue
		}

		common = collect(common, left[i])
		i++
		j++
	}

	return common
}
