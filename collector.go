// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

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
