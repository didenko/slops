// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

// Package slops (slice operations) is a collection of free functions
// operating on sorted slices of strings
package slops // import "didenko.com/go/slops"

// Common gathers same entries from two sorted slices into
// a new slice. The order is preserved. The lesser number of
// duplicates is preserved
func Common(left, right []string) []string {

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

		common = append(common, left[i])
		i++
		j++
	}

	return common
}
