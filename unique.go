// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Unique returns a new sorted slice without duplicates from
// the original sorted slice.
func Unique(in []string) []string {
	return CollectUnique(in, &getAll{})
}

// CollectUnique applies a Collector to every unique item in
// the sorted input slice.
func CollectUnique(in []string, c Collector) []string {

	uniq := make([]string, 0)

	if len(in) == 0 {
		return uniq
	}

	uniq = c.Collect(uniq, in[0])

	for i := 1; i < len(in); i++ {
		if in[i] != in[i-1] {
			uniq = c.Collect(uniq, in[i])
		}
	}

	return uniq
}
