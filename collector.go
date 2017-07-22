// Copyright 2017 Vlad Didenko. All rights reserved.
// See the included LICENSE.md file for licensing information

package slops // import "didenko.com/go/slops"

// Collector interface wraps the Collect method, so that an
// implementation can be used by other slops functions
type Collector interface {
	Collect(dest []string, item string) []string
}
