// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

// Common gathers same entries from two sorted slices into
// a new slice. The order is preserved. The lesser number of
// duplicates is preserved
func Common(left, right []string) []string {
	return CollectCommon(left, right, GetAll)
}

// CommonUnique gathers same entries from two sorted slices into
// a new slice. The order is preserved. Duplicates are reduced to
// a single item
func CommonUnique(left, right []string) []string {
	return CollectCommon(left, right, GetUnique)
}

// CollectCommon applies a Collector to every item which is
// in both left and right slices. Both input slices are
// expected to be sorted.
func CollectCommon(left, right []string, collect Collector) []string {
	return CollectVariety(left, right, NoOp, collect, NoOp)
}
