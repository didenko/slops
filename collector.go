// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

// Package slops (slice operations) is a collection of free functions
// operating on sorted slices of strings.
//
// It is designed around Collector functions, which act as folds.
// Once the basic logic (mostly implemented in `Collect...`
// functions) has determined a candidate slice element
// meeting a certain criteria, that element and the previously
// gathered slice are handed to a Collector function which
// decides how to modify the resulting slice conidering the new
// item.
//
// The expressive power comes from Collectors which close over
// some state or logic. For example, a Collector applied to
// differences of two file lists may open yet unopened files
// for further monitoring. Conveniently, the `ioutil.ReadDir`
// standard function returns files sorted.
//
// A good way to follow the suggested workflow is to see how the
// `MergeUnique` function uses `CollectVariety` with `getUnique`
// parameter.
package slops // import "go.didenko.com/slops"

// Collector is a type for function parameters so that
// implementations can be used by other slops functions
type Collector func(dest []string, item string) []string

func getAll(dest []string, item string) []string {
	return append(dest, item)
}

func getUnique(dest []string, item string) []string {
	if len(dest) > 0 && item == dest[len(dest)-1] {
		return dest
	}
	return append(dest, item)
}

func noop(dest []string, item string) []string {
	return dest
}
