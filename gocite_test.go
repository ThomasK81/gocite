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

//The testcorpora should get some more spelling names, according to the tasks they fulfill or possibly some documentation

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

//testcorpus2 is used to tests if sorting works with an empty passage in a work
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

//tescorpus5 is used to tests if sorting works with unordered passages in a work
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

var OldURNtests3 = []oldWorkTestgroup{ //slice format necessary?
	{inputCorpus: oldTestcorpus, inputID: "urn:cts:collection:workgroup.work:2-4", outputCorpus: oldTestcorpus2},
}

var URNtests3 = []newWorkTestgroup{ //slice format necessary?
	{inputCorpus: newTestcorpus, inputID: "urn:cts:collection:workgroup.work:2-4", outputCorpus: newTestcorpus2},
}

var oldURNtests4a = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus3, inputID: oldFirstPassageChange.First.PassageID},
	{inputCorpus: oldTestcorpus5, inputID: oldFirstPassageChange.First.PassageID},
}

var URNtests4a = []newWorkTestgroup{
	{inputCorpus: newTestcorpus3, inputID: newTestcorpus3.First.PassageID},
	{inputCorpus: newTestcorpus5, inputID: newTestcorpus5.First.PassageID},
}

var oldURNtests4 = []oldWorkTestgroup{
	{inputCorpus: oldTestcorpus2, outputCorpus: oldTestcorpus3},
	{inputCorpus: oldTestcorpus5, outputCorpus: oldTestcorpus6},
}

var URNtests4 = []newWorkTestgroup{
	{inputCorpus: newTestcorpus2, outputCorpus: newTestcorpus3},
	{inputCorpus: newTestcorpus5, outputCorpus: newTestcorpus6},
}

var oldTests5 = []oldWorkTestgroup{ //slice format necessary?
	{inputCorpus: oldTestcorpus7, outputCorpus: oldTestcorpus8},
}

var URNtests5 = []newWorkTestgroup{ //slice format necessary?
	{inputCorpus: newTestcorpus7, outputCorpus: newTestcorpus8},
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

func TestDePlassage(t *testing.T) {
	for _, workTestgroup := range URNtests3 {
		sortedWork, _ := gocite.DelPassage(workTestgroup.inputID, workTestgroup.inputCorpus)
		baseWork := oldTestWork{Ordered: sortedWork.Ordered}
		compareWork := oldTestWork{Ordered: workTestgroup.inputCorpus.Ordered}
		if baseWork.Ordered == compareWork.Ordered {
			t.Error(
				"For deleting", workTestgroup.inputID,
				"expected", !workTestgroup.inputCorpus.Ordered,
				"got", workTestgroup.inputCorpus.Ordered,
			)
		}
		for i := range sortedWork.Passages {
			basePassage := newTestPassage{PassageID: sortedWork.Passages[i].PassageID,
				Range: sortedWork.Passages[i].Range,
				Text:  sortedWork.Passages[i].Text,
				Index: sortedWork.Passages[i].Index,
				Prev:  sortedWork.Passages[i].Prev,
				Next:  sortedWork.Passages[i].Next}
			comparePassage := newTestPassage{PassageID: workTestgroup.outputCorpus.Passages[i].PassageID,
				Range: workTestgroup.outputCorpus.Passages[i].Range,
				Text:  workTestgroup.outputCorpus.Passages[i].Text,
				Index: workTestgroup.outputCorpus.Passages[i].Index,
				Prev:  workTestgroup.outputCorpus.Passages[i].Prev,
				Next:  workTestgroup.outputCorpus.Passages[i].Next}
			if basePassage != comparePassage {
				t.Error(
					"For deleting passage with ID ", workTestgroup.inputID,
					" expected remaining passage", workTestgroup.outputCorpus.Passages[i],
					", but got passage ", sortedWork.Passages[i], " instead",
				)
			}
		}
	}
}

func TestFindFirstIndex(t *testing.T) {
	for _, pair := range URNtests4a {
		v, found := gocite.FindFirstIndex(pair.inputCorpus)
		if found != true {
			t.Error(
				"For test ", pair.inputCorpus.WorkID,
				"expected", true,
				"got", false,
			)
		}
		if pair.inputCorpus.Passages[v].PassageID != pair.inputID {
			t.Error(
				"For test ", pair.inputCorpus.WorkID,
				"expected", pair.inputID,
				"got", pair.inputCorpus.Passages[v].PassageID,
			)
		}
	}
}

func TestSortPassages(t *testing.T) {
	/*for j, workTestgroup := range oldURNtests4 { //two runs
		var tempInputPassages []gocite.Passage
		for k := range workTestgroup.inputCorpus.Passages { //three runs
			var tempInputPassage = gocite.Passage{
				PassageID: workTestgroup.inputCorpus.Passages[k].PassageID,
				Range:     workTestgroup.inputCorpus.Passages[k].Range,
				Text:      workTestgroup.inputCorpus.Passages[k].Text,
				Index:     workTestgroup.inputCorpus.Passages[k].Index,
				Prev:      workTestgroup.inputCorpus.Passages[k].Prev,
				Next:      workTestgroup.inputCorpus.Passages[k].Next}
			tempInputPassages = append(tempInputPassages, tempInputPassage)
		}

		tempInPutCorpus := gocite.Work{
			WorkID:   workTestgroup.inputCorpus.WorkID,
			Passages: tempInputPassages,
			Ordered:  workTestgroup.inputCorpus.Ordered,
		}

		sortedWork, err := gocite.SortPassages(tempInPutCorpus)
		if sortedWork.Ordered == false {
			t.Error(
				"For sorting test ", j,
				" with release 1.0.1 data expected", true,
				" got ", false, ". Sorting failed.",
			)
		}
		if err != nil {
			t.Error("Error calling SortPassages: ", err)
		}

		for i := range sortedWork.Passages { //should run twice
			basePassage := newTestPassage{PassageID: sortedWork.Passages[i].PassageID,
				Range: sortedWork.Passages[i].Range,
				Text:  sortedWork.Passages[i].Text,
				Index: sortedWork.Passages[i].Index,
				Prev:  sortedWork.Passages[i].Prev,
				Next:  sortedWork.Passages[i].Next}
			comparePassage := newTestPassage{PassageID: workTestgroup.outputCorpus.Passages[i].PassageID,
				Range: workTestgroup.outputCorpus.Passages[i].Range,
				Text:  workTestgroup.outputCorpus.Passages[i].Text,
				Index: workTestgroup.outputCorpus.Passages[i].Index,
				Prev:  workTestgroup.outputCorpus.Passages[i].Prev,
				Next:  workTestgroup.outputCorpus.Passages[i].Next}
			if basePassage != comparePassage {
				t.Error(
					"For sorting test ", j, " passage ", i,
					" with release 1.0.1 data expected ", URNtests4[j].outputCorpus.Passages[i],
					" got ", sortedWork.Passages[i], " instead",
				)
			}
		}
	}*/

	for j, workTestgroup := range URNtests4 {
		work, err := gocite.SortPassages(workTestgroup.inputCorpus) //this could be adapted to account for the occurance of an error.
		if work.Ordered == false {
			t.Error(
				"For sorting test ", j,
				" with release 2.0.0 data expected", true,
				" got ", false, ". Sorting failed.",
			)
		}
		if err != nil {
			t.Error("Error calling SortPassages: ", err)
		}
		for i := range work.Passages {
			basePassage := newTestPassage{PassageID: work.Passages[i].PassageID,
				Range: work.Passages[i].Range,
				Text:  work.Passages[i].Text,
				Index: work.Passages[i].Index,
				Prev:  work.Passages[i].Prev,
				Next:  work.Passages[i].Next}
			comparePassage := newTestPassage{PassageID: workTestgroup.outputCorpus.Passages[i].PassageID,
				Range: workTestgroup.outputCorpus.Passages[i].Range,
				Text:  workTestgroup.outputCorpus.Passages[i].Text,
				Index: workTestgroup.outputCorpus.Passages[i].Index,
				Prev:  workTestgroup.outputCorpus.Passages[i].Prev,
				Next:  workTestgroup.outputCorpus.Passages[i].Next}
			if basePassage != comparePassage {
				t.Error(
					"For sorting test ", j, " passage ", i,
					" with release 2.0.0 data expected ", workTestgroup.outputCorpus.Passages[i],
					" got ", work.Passages[i], " instead",
				)
			}
		}
	}
}

func TestInsertPassage(t *testing.T) {
	v, err := gocite.InsertPassage(newSecondPassage, URNtests5[0].inputCorpus)
	if v.Ordered == true {
		t.Error(
			"Expected ordered", false,
			"got", true,
		)
	}
	if err != nil {
		t.Error(
			"Error calling InsertPassage: ", err,
		)
	}
	for i := range v.Passages {
		basePassage := oldTestPassage{PassageID: v.Passages[i].PassageID,
			Range: v.Passages[i].Range,
			Text:  v.Passages[i].Text,
			Index: v.Passages[i].Index,
			Prev:  v.Passages[i].Prev,
			Next:  v.Passages[i].Next}
		comparePassage := oldTestPassage{PassageID: URNtests5[0].outputCorpus.Passages[i].PassageID,
			Range: URNtests5[0].outputCorpus.Passages[i].Range,
			Text:  URNtests5[0].outputCorpus.Passages[i].Text,
			Index: URNtests5[0].outputCorpus.Passages[i].Index,
			Prev:  URNtests5[0].outputCorpus.Passages[i].Prev,
			Next:  URNtests5[0].outputCorpus.Passages[i].Next}
		if basePassage != comparePassage {
			t.Error(
				"For test", i,
				"expected", URNtests5[0].outputCorpus.Passages[i],
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

var oldPassageOne = oldTestPassage{
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

var PassageOne = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:1",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is is the first node.",
	},
	Index: 0,
	Prev:  gocite.PassLoc{},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2", Index: 1},
}

var oldPassageTwo = oldTestPassage{
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

var PassageTwo = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:2",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is. the second. node.",
	},
	Index: 1,
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:1", Index: 0},
	Next:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:3", Index: 2},
}

var oldPassageThree = oldTestPassage{
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

var PassageThree = gocite.Passage{
	PassageID: "urn:cts:collection:workgroup.work:3",
	Range:     false,
	Text: gocite.EncText{
		TXT: "This is the third node.",
	},
	Index: 2,
	Prev:  gocite.PassLoc{Exists: true, PassageID: "urn:cts:collection:workgroup.work:2", Index: 1},
	Next:  gocite.PassLoc{},
}
