// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Exclude returns a new slice where strings in the rejects slice
// removed from the src slice. Both slices expected to be sorted.
func Exclude(src, rejects []string) []string {
	return CollectExcluded(src, rejects, getAll)
}

// CollectExcluded applies a Collector to every item which is in
// src slice but not in the rejects slice. Both input slices are
// expected to be sorted.
func CollectExcluded(src, rejects []string, collect Collector) []string {

	filtered := make([]string, 0)

	for i, j := 0, 0; i < len(src); {

		if j >= len(rejects) || src[i] < rejects[j] {
			filtered = collect(filtered, src[i])
			i++
			continue
		}

		if src[i] > rejects[j] {
			j++
			continue
		}

		// last option is src[i] == rejects[j]
		i++
	}

	return filtered
}
