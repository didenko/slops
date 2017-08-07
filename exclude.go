// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

// ExcludeAll returns a new slice where all strings from the
// rejects slice are removed from the src slice, regardless of
// how many times they occur in either slice. Non-excluded
// duplicates in the src slice are preserved. Both slices
// are expected to be sorted.
func ExcludeAll(src, rejects []string) []string {
	return CollectExcludedAll(src, rejects, GetAll)
}

// CollectExcludedAll applies a Collector to every item which is
// in src slice but not in the rejects slice. All strings from
// the rejects slice are removed from the src slice, regardless
// of how many times they occur in either slice. Non-excluded
// duplicates in the src slice are preserved. Both input slices
// are expected to be sorted.
func CollectExcludedAll(src, rejects []string, collect Collector) []string {

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

// CollectExcludedByCount applies a Collector to every item which
// is in src slice but not in the rejects slice. Only as many
// duplicates removed from the src slice as occur in the rejects
// slice. Both input slices are expected to be sorted.
func CollectExcludedByCount(src, rejects []string, collect Collector) []string {
	return CollectVariety(src, rejects, collect, NoOp, NoOp)
}

// ExcludeAllContrived is a companion function to
// CollectExcludedAllContrived. This is a nerdy exersize
// and is not intended for production use.
func ExcludeAllContrived(src, rejects []string) []string {
	return CollectExcludedAllContrived(src, rejects, GetAll)
}

// CollectExcludedAllContrived has exactly the same functionality
// as CollectExcludedAll function (and is tested against the same
// use cases) but is slower and is presented to demonstrate how
// similar effects can be achieved from the generalised function
// CollectVariety via functional programming means. This is merely
// a nerdy exersize and is not intended for production use.
func CollectExcludedAllContrived(src, rejects []string, collect Collector) []string {
	var lastExcluded string
	return CollectVariety(
		src,
		rejects,
		func(dest []string, item string) []string {
			if item != lastExcluded {
				return collect(dest, item)
			}
			return dest
		},
		func(dest []string, item string) []string {
			lastExcluded = item
			return dest
		},
		NoOp)
}
