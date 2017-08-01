// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Merge returns a slice with a union of strings in slices.
// For duplicate entries, the resulting slice contains the
// maximum numbers of duplicate strings between the original
// slices. Both left and right slices are expected to be sorted.
func Merge(left, right []string) []string {
	return CollectVariety(left, right, getAll, getAll, getAll)
}

// MergeUnique returns a slice with a union of strings in slices.
// The resulting slice contains one entry for each set  of
// duplicate strings in the original slices. Both left
// and right slices are expected to be sorted.
func MergeUnique(left, right []string) []string {
	return CollectVariety(left, right, getUnique, getUnique, getUnique)
}
