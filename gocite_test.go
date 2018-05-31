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

type testpair2 struct {
	input                                        string
	outputTG, outputWork, outputVers, outputExmp bool
}

type testgroup struct {
	inputcorpus gocite.Work
	inputID     string
	output      gocite.Work
}

var tests = []testpair{
	{input: "urn:cts:collection:workgroup.work:1-27", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:1-27", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "1-27"}, outputRange: true, outputCTS: true},
	{input: "urn:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:27.3", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "27.3"}, outputRange: false, outputCTS: true},
	{input: "not:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "not:cts:collection:workgroup.work:27.3", InValid: true}, outputRange: false, outputCTS: false}}

var tests2 = []testpair2{
	{input: "urn:cts:collection:workgroup:", outputTG: true, outputWork: false, outputVers: false, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work:27.3", outputTG: false, outputWork: true, outputVers: false, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work.version:27.3-29.9", outputTG: false, outputWork: false, outputVers: true, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work.version.exemplar:29.9", outputTG: false, outputWork: false, outputVers: false, outputExmp: true},
}

var firstPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the first node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Prev:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var firstPassageChange = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the first node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Prev:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
}

var thirdPassageChange = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
}

var secondPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:2-4",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is. the second. node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
}

var thirdPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var testcorpus = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		firstPassage,
		secondPassage,
		thirdPassage,
	},
	Ordered: true,
}

var testcorpus2 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		firstPassageChange,
		gocite.Passage{},
		thirdPassageChange,
	},
	Ordered: false}

var testcorpus3 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		firstPassageChange,
		thirdPassageChange,
	},
	Ordered: true}

var testcorpus4 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		firstPassageChange,
		thirdPassageChange,
	},
	Ordered: true}

var testcorpus5 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		gocite.Passage{PassageID: firstPassageChange.PassageID, Range: firstPassageChange.Range, Text: firstPassageChange.Text, Index: 0, First: firstPassageChange.First, Last: firstPassageChange.Last, Prev: firstPassageChange.Prev, Next: gocite.PassLoc{Exists: true, PassageID: secondPassage.PassageID, Index: 2}},
		gocite.Passage{PassageID: thirdPassageChange.PassageID, Range: thirdPassageChange.Range, Text: thirdPassageChange.Text, Index: 1, First: firstPassageChange.First, Last: firstPassageChange.Last, Prev: gocite.PassLoc{Exists: true, PassageID: secondPassage.PassageID, Index: 2}, Next: gocite.PassLoc{}},
		gocite.Passage{PassageID: secondPassage.PassageID, Range: secondPassage.Range, Text: secondPassage.Text, Index: 2, First: thirdPassageChange.First, Last: thirdPassageChange.Last, Prev: thirdPassageChange.First, Next: gocite.PassLoc{Exists: true, PassageID: thirdPassage.PassageID, Index: 1}},
	},
	Ordered: false,
}

var testcorpus6 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		gocite.Passage{PassageID: firstPassageChange.PassageID, Range: firstPassageChange.Range, Text: firstPassageChange.Text, Index: 0, First: gocite.PassLoc{Exists: true, PassageID: firstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: thirdPassageChange.PassageID, Index: 2}, Next: gocite.PassLoc{Exists: true, PassageID: secondPassage.PassageID, Index: 1}},
		gocite.Passage{PassageID: secondPassage.PassageID, Range: secondPassage.Range, Text: secondPassage.Text, Index: 1, First: gocite.PassLoc{Exists: true, PassageID: firstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: thirdPassageChange.PassageID, Index: 2}, Next: gocite.PassLoc{Exists: true, PassageID: thirdPassageChange.PassageID, Index: 2}, Prev: gocite.PassLoc{Exists: true, PassageID: firstPassageChange.PassageID, Index: 0}},
		gocite.Passage{PassageID: thirdPassageChange.PassageID, Range: thirdPassageChange.Range, Text: thirdPassageChange.Text, Index: 2, First: gocite.PassLoc{Exists: true, PassageID: firstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: thirdPassageChange.PassageID, Index: 2}, Prev: gocite.PassLoc{Exists: true, PassageID: secondPassage.PassageID, Index: 0}},
	},
	Ordered: false,
}

var tests3 = []testgroup{
	testgroup{inputcorpus: testcorpus, inputID: "urn:cts:collection:workgroup.work:2-4", output: testcorpus2},
}

var tests4a = []testgroup{
	testgroup{inputcorpus: testcorpus3, inputID: firstPassageChange.First.PassageID},
	testgroup{inputcorpus: testcorpus5, inputID: firstPassageChange.First.PassageID},
}

var tests4 = []testgroup{
	testgroup{inputcorpus: testcorpus2, output: testcorpus3},
	testgroup{inputcorpus: testcorpus5, output: testcorpus6},
}

var tests5 = []testgroup{
	testgroup{inputcorpus: testcorpus4, output: testcorpus5},
}

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

func TestIsRange(t *testing.T) {
	for _, pair := range tests {
		v := gocite.IsRange(pair.input)
		if v != pair.outputRange {
			t.Error(
				"For", pair.input,
				"expected", pair.outputRange,
				"got", v,
			)
		}
	}
}

func TestIsCTSURN(t *testing.T) {
	for _, pair := range tests {
		v := gocite.IsCTSURN(pair.input)
		if v != pair.outputCTS {
			t.Error(
				"For", pair.input,
				"expected", pair.outputCTS,
				"got", v,
			)
		}
	}
}

func TestIsTextgroupID(t *testing.T) {
	for _, pair := range tests2 {
		v := gocite.IsTextgroupID(pair.input)
		if v != pair.outputTG {
			t.Error(
				"For", pair.input,
				"expected", pair.outputTG,
				"got", v,
			)
		}
	}
}

func TestIsWorkID(t *testing.T) {
	for _, pair := range tests2 {
		v := gocite.IsWorkID(pair.input)
		if v != pair.outputWork {
			t.Error(
				"For", pair.input,
				"expected", pair.outputWork,
				"got", v,
			)
		}
	}
}

func TestIsVersionID(t *testing.T) {
	for _, pair := range tests2 {
		v := gocite.IsVersionID(pair.input)
		if v != pair.outputVers {
			t.Error(
				"For", pair.input,
				"expected", pair.outputVers,
				"got", v,
			)
		}
	}
}

func TestIsExemplarID(t *testing.T) {
	for _, pair := range tests2 {
		v := gocite.IsExemplarID(pair.input)
		if v != pair.outputExmp {
			t.Error(
				"For", pair.input,
				"expected", pair.outputExmp,
				"got", v,
			)
		}
	}
}

func TestDelPassage(t *testing.T) {
	for _, pair := range tests3 {
		v := gocite.DelPassage(pair.inputID, pair.inputcorpus)
		if v.Ordered == pair.inputcorpus.Ordered {
			t.Error(
				"For deleting", pair.inputID,
				"expected", !pair.inputcorpus.Ordered,
				"got", pair.inputcorpus.Ordered,
			)
		}
		for i := range v.Passages {
			if v.Passages[i] != pair.output.Passages[i] {
				t.Error(
					"For deleting", pair.inputID,
					"expected", pair.output.Passages[i],
					"got", v.Passages[i],
				)
			}
		}
	}
}

func TestFindFirstIndex(t *testing.T) {
	for _, pair := range tests4a {
		v, found := gocite.FindFirstIndex(pair.inputcorpus)
		if found != true {
			t.Error(
				"For test ", pair.inputcorpus.WorkID,
				"expected", true,
				"got", false,
			)
		}
		if pair.inputcorpus.Passages[v].PassageID != pair.inputID {
			t.Error(
				"For test ", pair.inputcorpus.WorkID,
				"expected", pair.inputID,
				"got", pair.inputcorpus.Passages[v].PassageID,
			)
		}
	}
}

func TestSortPassages(t *testing.T) {
	for j, pair := range tests4 {
		v := gocite.SortPassages(pair.inputcorpus)
		if v.Ordered == false {
			t.Error(
				"For test ", j,
				"expected", true,
				"got", false,
			)
		}
		for i := range v.Passages {
			if v.Passages[i] != pair.output.Passages[i] {
				t.Error(
					"For test ", j, i,
					"expected", pair.output.Passages[i],
					"got", v.Passages[i],
				)
			}
		}
	}
}

func TestInsertPassage(t *testing.T) {
	v := gocite.InsertPassage(secondPassage, tests5[0].inputcorpus)
	if v.Ordered == true {
		t.Error(
			"Expected ordered", false,
			"got", true,
		)
	}
	for i := range v.Passages {
		if v.Passages[i] != tests5[0].output.Passages[i] {
			t.Error(
				"For test ", i,
				"expected", tests5[0].output.Passages[i],
				"got", v.Passages[i],
			)
		}
	}
}
