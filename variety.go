// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "go.didenko.com/slops"

// CollectVariety applies a Collector to every item in left and
// right slices. Common items are collected as many times as
// there are common occurences by the related collector. Non-common
// items each, with their duplicates which are extra to the common
// items, are collected by related leftOnly and rightOnly Collectors.
// Both left and right input slices are expected to be sorted.
func CollectVariety(left, right []string, onlyLeft, common, onlyRight Collector) []string {

	collected := make([]string, 0)

	for l, r := 0, 0; l < len(left) || r < len(right); {

		if l < len(left) && (r == len(right) || left[l] < right[r]) {
			collected = onlyLeft(collected, left[l])
			l++
			continue
		}

		if r < len(right) && (l == len(left) || left[l] > right[r]) {
			collected = onlyRight(collected, right[r])
			r++
			continue
		}

		collected = common(collected, left[l])
		l++
		r++
	}

	return collected
}
