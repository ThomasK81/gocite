package gocite_test

import (
	"testing"

	"github.com/ThomasK81/gocite"
)

type oldTestPassage struct {
	PassageID               string
	Range                   bool
	Text                    gocite.EncText
	Index                   int
	First, Last, Prev, Next gocite.PassLoc
}

type newTestPassage struct {
	PassageID  string
	Range      bool
	Text       gocite.EncText
	Index      int
	Prev, Next gocite.PassLoc
}

type oldTestWork struct {
	WorkID   string
	Passages []oldTestPassage
	Ordered  bool
}

type newTestWork struct {
	WorkID      string
	Passages    []oldTestPassage
	Ordered     bool
	First, Last gocite.PassLoc
}

type URNTestpair struct {
	input                  string
	outputSplit            gocite.CTSURN
	outputRange, outputCTS bool
}

type testIDfeaturesTestpair struct {
	input                                        string
	outputTG, outputWork, outputVers, outputExmp bool
}

type oldWorkTestgroup struct {
	inputCorpus  oldTestWork
	inputID      string
	outputCorpus oldTestWork
}

type newWorkTestgroup struct {
	inputCorpus  gocite.Work
	inputID      string
	outputCorpus gocite.Work
}

type extractgroup struct {
	input  string
	answer []gocite.TextAndID
}

var URNtests = []URNTestpair{
	{input: "urn:cts:collection:workgroup.work:1-27", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:1-27", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "1-27"}, outputRange: true, outputCTS: true},
	{input: "urn:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "urn:cts:collection:workgroup.work:27.3", Base: "urn", Protocol: "cts", Namespace: "collection", Work: "workgroup.work", Passage: "27.3"}, outputRange: false, outputCTS: true},
	{input: "not:cts:collection:workgroup.work:27.3", outputSplit: gocite.CTSURN{ID: "not:cts:collection:workgroup.work:27.3", InValid: true}, outputRange: false, outputCTS: false}}

var IDfeatureTests = []testIDfeaturesTestpair{
	{input: "urn:cts:collection:workgroup:", outputTG: true, outputWork: false, outputVers: false, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work:27.3", outputTG: false, outputWork: true, outputVers: false, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work.version:27.3-29.9", outputTG: false, outputWork: false, outputVers: true, outputExmp: false},
	{input: "urn:cts:collection:workgroup.work.version.exemplar:29.9", outputTG: false, outputWork: false, outputVers: false, outputExmp: true},
}

var oldFirstPassage = oldTestPassage{
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

var newFirstPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the first node.",
	},
	Index: 0,
	Prev:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var oldFirstPassageChange = oldTestPassage{
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

var newFirstPassageChange = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the first node.",
	},
	Index: 0,
	Prev:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
}

var oldSecondPassage = oldTestPassage{
	PassageID: "urn:cts:collection:workgroup.work:2-4",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is. the second. node.",
	},
	Index: 1,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
}

var newSecondPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:2-4",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is. the second. node.",
	},
	Index: 1,
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
}

var oldThirdPassage = oldTestPassage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var newThirdPassage = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2-4", Index: 1},
}

var oldThirdPassageChange = oldTestPassage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
}

var newThirdPassageChange = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:5",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	Next:  gocite.PassLoc{Exists: false, PassageID: "", Index: 0},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
}

var oldTestcorpus = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		oldFirstPassage,
		oldSecondPassage,
		oldThirdPassage,
	},
	Ordered: true,
}

var newTestcorpus = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		newFirstPassage,
		newSecondPassage,
		newThirdPassage,
	},
	First:   gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Ordered: true,
}

var oldTestcorpus2 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		oldFirstPassageChange,
		{},
		oldThirdPassageChange,
	},
	Ordered: false}

var newTestcorpus2 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		newFirstPassageChange,
		{},
		newThirdPassageChange,
	},
	First:   gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:5", Index: 2},
	Ordered: false}

var oldTestcorpus3 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		{PassageID: oldFirstPassageChange.PassageID, Range: oldFirstPassageChange.Range, Text: oldFirstPassageChange.Text, Index: 0, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}},
		{PassageID: oldThirdPassageChange.PassageID, Range: oldThirdPassageChange.Range, Text: oldThirdPassageChange.Text, Index: 1, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Prev: gocite.PassLoc{Exists: true, PassageID: oldFirstPassage.PassageID, Index: 0}},
	},
	Ordered: true}

var newTestcorpus3 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		{PassageID: newFirstPassageChange.PassageID, Range: newFirstPassageChange.Range, Text: newFirstPassageChange.Text, Index: 0, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}},
		{PassageID: newThirdPassageChange.PassageID, Range: newThirdPassageChange.Range, Text: newThirdPassageChange.Text, Index: 1, Prev: gocite.PassLoc{Exists: true, PassageID: oldFirstPassage.PassageID, Index: 0}},
	},
	First:   gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1},
	Ordered: true}

var oldTestcorpus4 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		oldFirstPassageChange,
		oldThirdPassageChange,
	},
	Ordered: true}

var newTestcorpus4 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		newFirstPassageChange,
		newThirdPassageChange,
	},
	First:   gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1},
	Ordered: true}

var oldTestcorpus5 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		{PassageID: oldFirstPassageChange.PassageID, Range: oldFirstPassageChange.Range, Text: oldFirstPassageChange.Text, Index: 0, First: oldFirstPassageChange.First, Last: oldFirstPassageChange.Last, Prev: oldFirstPassageChange.Prev, Next: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 2}},
		{PassageID: oldThirdPassageChange.PassageID, Range: oldThirdPassageChange.Range, Text: oldThirdPassageChange.Text, Index: 1, First: oldFirstPassageChange.First, Last: oldFirstPassageChange.Last, Prev: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 2}, Next: gocite.PassLoc{}},
		{PassageID: oldSecondPassage.PassageID, Range: oldSecondPassage.Range, Text: oldSecondPassage.Text, Index: 2, First: oldThirdPassageChange.First, Last: oldThirdPassageChange.Last, Prev: oldThirdPassageChange.First, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassage.PassageID, Index: 1}},
	},
	Ordered: false,
}

var newTestcorpus5 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		{PassageID: newFirstPassageChange.PassageID, Range: newFirstPassageChange.Range, Text: newFirstPassageChange.Text, Index: 0, Prev: newFirstPassageChange.Prev, Next: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 2}},
		{PassageID: newThirdPassageChange.PassageID, Range: newThirdPassageChange.Range, Text: newThirdPassageChange.Text, Index: 1, Prev: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 2}, Next: gocite.PassLoc{}},
		{PassageID: newSecondPassage.PassageID, Range: newSecondPassage.Range, Text: newSecondPassage.Text, Index: 2, Prev: gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 0}, Next: gocite.PassLoc{Exists: true, PassageID: newThirdPassage.PassageID, Index: 1}},
	},
	First:   oldFirstPassageChange.First,
	Last:    oldFirstPassageChange.Last,
	Ordered: false,
}

var oldTestcorpus6 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		{PassageID: oldFirstPassageChange.PassageID, Range: oldFirstPassageChange.Range, Text: oldFirstPassageChange.Text, Index: 0, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 2}, Next: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 1}},
		{PassageID: oldSecondPassage.PassageID, Range: oldSecondPassage.Range, Text: oldSecondPassage.Text, Index: 1, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 2}, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 2}, Prev: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}},
		{PassageID: oldThirdPassageChange.PassageID, Range: oldThirdPassageChange.Range, Text: oldThirdPassageChange.Text, Index: 2, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 2}, Prev: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 1}},
	},
	Ordered: false,
}

var newTestcorpus6 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		{PassageID: newFirstPassageChange.PassageID, Range: newFirstPassageChange.Range, Text: newFirstPassageChange.Text, Index: 0, Next: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 1}},
		{PassageID: newSecondPassage.PassageID, Range: newSecondPassage.Range, Text: newSecondPassage.Text, Index: 1, Next: gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 2}, Prev: gocite.PassLoc{Exists: true, PassageID: newFirstPassageChange.PassageID, Index: 0}},
		{PassageID: newThirdPassageChange.PassageID, Range: newThirdPassageChange.Range, Text: newThirdPassageChange.Text, Index: 2, Prev: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 1}},
	},
	First:   gocite.PassLoc{Exists: true, PassageID: newFirstPassageChange.PassageID, Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 2},
	Ordered: false,
}

var oldTestcorpus7 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		{PassageID: oldFirstPassageChange.PassageID, Range: false, Text: oldFirstPassageChange.Text, Index: 0, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassage.PassageID, Index: 1}},
		{PassageID: oldThirdPassageChange.PassageID, Range: false, Text: oldThirdPassageChange.Text, Index: 1, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Prev: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}},
	},
	Ordered: true}

var newTestcorpus7 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		{PassageID: newFirstPassageChange.PassageID, Range: false, Text: newFirstPassageChange.Text, Index: 0, Next: gocite.PassLoc{Exists: true, PassageID: newThirdPassage.PassageID, Index: 1}},
		{PassageID: newThirdPassageChange.PassageID, Range: false, Text: newThirdPassageChange.Text, Index: 1, Prev: gocite.PassLoc{Exists: true, PassageID: newFirstPassageChange.PassageID, Index: 0}},
	},
	First:   gocite.PassLoc{Exists: true, PassageID: newFirstPassageChange.PassageID, Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 1},
	Ordered: true}

var oldTestcorpus8 = oldTestWork{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []oldTestPassage{
		{PassageID: oldFirstPassageChange.PassageID, Range: false, Text: oldFirstPassageChange.Text, Index: 0, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Next: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 2}},
		{PassageID: oldThirdPassageChange.PassageID, Range: false, Text: oldThirdPassageChange.Text, Index: 1, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Prev: gocite.PassLoc{Exists: true, PassageID: oldSecondPassage.PassageID, Index: 2}},
		{PassageID: oldSecondPassage.PassageID, Range: oldSecondPassage.Range, Text: oldSecondPassage.Text, Index: 2, First: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Last: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}, Prev: gocite.PassLoc{Exists: true, PassageID: oldFirstPassageChange.PassageID, Index: 0}, Next: gocite.PassLoc{Exists: true, PassageID: oldThirdPassageChange.PassageID, Index: 1}},
	},
	Ordered: false,
}

var newTestcorpus8 = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		{PassageID: newFirstPassageChange.PassageID, Range: false, Text: newFirstPassageChange.Text, Index: 0, Next: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 2}},
		{PassageID: newThirdPassageChange.PassageID, Range: false, Text: newThirdPassageChange.Text, Index: 1, Prev: gocite.PassLoc{Exists: true, PassageID: newSecondPassage.PassageID, Index: 2}},
		{PassageID: newSecondPassage.PassageID, Range: newSecondPassage.Range, Text: newSecondPassage.Text, Index: 2, Next: gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 1}},
	},

	First:   gocite.PassLoc{Exists: true, PassageID: newFirstPassageChange.PassageID, Index: 0},
	Last:    gocite.PassLoc{Exists: true, PassageID: newThirdPassageChange.PassageID, Index: 1},
	Ordered: false,
}

var URNtests3 = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus, inputID: "urn:cts:collection:workgroup.work:2-4", outputCorpus: oldTestcorpus2},
}

var URNtests4a = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus3, inputID: oldFirstPassageChange.First.PassageID},
	{inputCorpus: oldTestcorpus5, inputID: oldFirstPassageChange.First.PassageID},
}

var URNtests4 = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus2, outputCorpus: oldTestcorpus3},
	{inputCorpus: oldTestcorpus5, outputCorpus: oldTestcorpus6},
}

var URNtests5 = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus7, outputCorpus: oldTestcorpus8},
}

func TestSplitCTS(t *testing.T) {
	for _, pair := range URNtests {
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
	for _, pair := range URNtests {
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
	for _, pair := range URNtests {
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
	for _, pair := range IDfeatureTests {
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
	for _, pair := range IDfeatureTests {
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
	for _, pair := range IDfeatureTests {
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
	for _, pair := range IDfeatureTests {
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
	for _, pair := range URNtests3 {
		v, _ := gocite.DelPassage(pair.inputID, pair.inputcorpus)
		baseWork := oldTestWork{Ordered: v.Ordered}
		compareWork := oldTestWork{Ordered: pair.inputcorpus.Ordered}
		if baseWork.Ordered == compareWork.Ordered {
			t.Error(
				"For deleting", pair.inputID,
				"expected", !pair.inputcorpus.Ordered,
				"got", pair.inputcorpus.Ordered,
			)
		}
		for i := range v.Passages {
			basePassage := oldTestPassage{PassageID: v.Passages[i].PassageID,
				Range: v.Passages[i].Range,
				Text:  v.Passages[i].Text,
				Index: v.Passages[i].Index,
				First: v.Passages[i].First,
				Last:  v.Passages[i].Last,
				Prev:  v.Passages[i].Prev,
				Next:  v.Passages[i].Next}
			comparePassage := oldTestPassage{PassageID: pair.output.Passages[i].PassageID,
				Range: pair.output.Passages[i].Range,
				Text:  pair.output.Passages[i].Text,
				Index: pair.output.Passages[i].Index,
				First: pair.output.Passages[i].First,
				Last:  pair.output.Passages[i].Last,
				Prev:  pair.output.Passages[i].Prev,
				Next:  pair.output.Passages[i].Next}
			if basePassage != comparePassage {
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
	for _, pair := range URNtests4a {
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
	for j, pair := range URNtests4 {
		v := gocite.SortPassages(pair.inputcorpus)
		if v.Ordered == false {
			t.Error(
				"For test ", j,
				"expected", true,
				"got", false,
			)
		}
		for i := range v.Passages {
			basePassage := oldTestPassage{PassageID: v.Passages[i].PassageID,
				Range: v.Passages[i].Range,
				Text:  v.Passages[i].Text,
				Index: v.Passages[i].Index,
				First: v.Passages[i].First,
				Last:  v.Passages[i].Last,
				Prev:  v.Passages[i].Prev,
				Next:  v.Passages[i].Next}
			comparePassage := oldTestPassage{PassageID: pair.output.Passages[i].PassageID,
				Range: pair.output.Passages[i].Range,
				Text:  pair.output.Passages[i].Text,
				Index: pair.output.Passages[i].Index,
				First: pair.output.Passages[i].First,
				Last:  pair.output.Passages[i].Last,
				Prev:  pair.output.Passages[i].Prev,
				Next:  pair.output.Passages[i].Next}
			if basePassage != comparePassage {
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
	v := gocite.InsertPassage(oldSecondPassage, URNtests5[0].inputcorpus)
	if v.Ordered == true {
		t.Error(
			"Expected ordered", false,
			"got", true,
		)
	}
	for i := range v.Passages {
		basePassage := oldTestPassage{PassageID: v.Passages[i].PassageID,
			Range: v.Passages[i].Range,
			Text:  v.Passages[i].Text,
			Index: v.Passages[i].Index,
			First: v.Passages[i].First,
			Last:  v.Passages[i].Last,
			Prev:  v.Passages[i].Prev,
			Next:  v.Passages[i].Next}
		comparePassage := oldTestPassage{PassageID: URNtests5[0].output.Passages[i].PassageID,
			Range: URNtests5[0].output.Passages[i].Range,
			Text:  URNtests5[0].output.Passages[i].Text,
			Index: URNtests5[0].output.Passages[i].Index,
			First: URNtests5[0].output.Passages[i].First,
			Last:  URNtests5[0].output.Passages[i].Last,
			Prev:  URNtests5[0].output.Passages[i].Prev,
			Next:  URNtests5[0].output.Passages[i].Next}
		if basePassage != comparePassage {
			t.Error(
				"For test", i,
				"expected", URNtests5[0].output.Passages[i],
				"got", v.Passages[i],
			)
		}
	}
}

func TestExtractTextByID(t *testing.T) {
	for j := range extrtest {
		request := extrtest[j].input
		answer := extrtest[j].answer
		v, err := gocite.ExtractTextByID(request, testExtrcorpus)
		if err != nil {
			t.Error(request, err)
		}
		if len(v) != len(answer) {
			t.Error(
				"For test", request,
				"expected", len(answer),
				"got", len(v),
			)
		}
		for i := range v {
			if v[i] != answer[i] {
				t.Error(
					"For test", request, "part", i,
					"expected", answer,
					"got", v,
				)
			}
		}
	}
}

var extrtest = []extractgroup{
	{
		input: "urn:cts:collection:workgroup.work:1",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1", Text: "This is is the first node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:2",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1-3",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1", Text: "This is is the first node."},
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3", Text: "This is the third node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:2-3",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3", Text: "This is the third node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is", Text: "is"},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is[2]",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is[2]", Text: "is"},
		}},
	{
		input: "urn:cts:collection:workgroup.work:2@is[2]-2@second",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:2@is[2]-2@second", Text: "is. the second"},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is-3",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is", Text: "is is is the first node."},
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3", Text: "This is the third node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is[2]-3",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is[2]", Text: "is is the first node."},
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3", Text: "This is the third node."},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is[2]-3@third",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is[2]", Text: "is is the first node."},
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3@third", Text: "This is the third"},
		}},
	{
		input: "urn:cts:collection:workgroup.work:1@is[2]-3@is[2]",
		answer: []gocite.TextAndID{
			{ID: "urn:cts:collection:workgroup.work:1@is[2]", Text: "is is the first node."},
			{ID: "urn:cts:collection:workgroup.work:2", Text: "This is. the second. node."},
			{ID: "urn:cts:collection:workgroup.work:3@is[2]", Text: "This is"},
		}},
}

var testExtrcorpus = gocite.Work{
	WorkID: "urn:cts:collection:workgroup.work:",
	Passages: []gocite.Passage{
		PassageOne,
		PassageTwo,
		PassageThree,
	},
	Ordered: true,
}

var PassageOne = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is is the first node.",
	},
	Index: 0,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:3", Index: 2},
	Prev:  gocite.PassLoc{},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2", Index: 1},
}

var PassageTwo = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:2",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is. the second. node.",
	},
	Index: 1,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:3", Index: 2},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:3", Index: 2},
}

var PassageThree = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:3",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	First: gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Last:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:3", Index: 2},
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2", Index: 1},
	Next:  gocite.PassLoc{},
}
