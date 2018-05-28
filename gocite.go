package gocite

import (
	"strings"
)

// URN references text as explained in http://cite-architecture.github.io
type URN struct {
	Stem      string
	Reference string
}

// Work is an ordered container for CTS passages that belong to the same work
type Work struct {
	WorkURN string
	URN     []string
	Text    []string
	Index   []int
}

// Workgroup is a collection of works
type Workgroup struct {
	URN   []string
	Works []Work
}

// SplitCTS splits a CTS URN in its stem and the passage reference
func SplitCTS(s string) URN {
	var result URN
	result = URN{Stem: strings.Join(strings.Split(s, ":")[0:4], ":") + ":", Reference: strings.Split(s, ":")[4]}
	return result
}

// IsRange is a function that returns a boolean whether a CTS URN is a range
func IsRange(s string) bool {
	switch {
	case len(strings.Split(s, ":")) < 5:
		return false
	case strings.Contains(strings.Split(s, ":")[4], "-"):
		return true
	default:
		return false
	}
}

// Boolean function whether a string is a valid CTS URN
func isCTSURN(s string) bool {
	test := strings.Split(s, ":")
	switch {
	case len(test) < 4:
		return false
	case len(test) > 5:
		return false
	case test[0] != "urn":
		return false
	case test[1] != "cts":
		return false
	default:
		return true
	}
}
