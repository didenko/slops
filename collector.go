// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Collector interface wraps the Collect method, so that an
// implementation can be used by other slops functions
type Collector interface {
	Collect(dest []string, item string) []string
}

type getAll struct{}

func (ga *getAll) Collect(dest []string, item string) []string {
	return append(dest, item)
}

type getUnique struct{}

func (gu *getUnique) Collect(dest []string, item string) []string {
	if len(dest) > 0 && item == dest[len(dest)-1] {
		return dest
	}
	return append(dest, item)
}
