package gocite_test

import (
	"testing"

	"github.com/ThomasK81/gocite"
)

type testpair struct {
	input                  string
	outputSplit            gocite.CTSURN
	outputRange, outputCTS bool
}

var tests = []testpair{
	{input: "urn:cts:collection:workgroup.work:1-27", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:1-27", Base: "urn", Protocol: "cts", Namespace: "workgroup", Work: "work", Passage: "1-27"}, outputRange: true, outputCTS: true},
	{input: "urn:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:27.3", Base: "urn", Protocol: "cts", Namespace: "workgroup", Work: "work", Passage: "27.3"}, outputRange: false, outputCTS: true},
	{input: "not:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "not:cts:collection:workgroup.work:27.3", InValid: true}, outputRange: false, outputCTS: false}}

func TestSplitCTS(t *testing.T) {
	for _, pair := range tests {
		v := gocite.SplitCTS(pair.input)
		if v != pair.outputSplit {
			t.Error(
				"For", pair.input,
				"expected", pair.outputSplit,
				"got", v,
			)
		}
	}
}
