package gocite_test

import (
	"testing"

	"github.com/ThomasK81/gocite"
)

type testpair struct {
	input  string
	output gocite.URN
}

var tests = []testpair{
	{"urn:cts:collection:workgroup.work:1-27", gocite.URN{Stem: "urn:cts:collection:workgroup.work:", Reference: "1-27"}},
	{"urn:cts:collection:workgroup.work:27.3", gocite.URN{Stem: "urn:cts:collection:workgroup.work:", Reference: "27.3"}},
}

func TestSplitCTS(t *testing.T) {
	for _, pair := range tests {
		v := gocite.SplitCTS(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}
