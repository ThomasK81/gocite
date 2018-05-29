package gocite

import (
	"strings"
)

// CTSURN references text as explained in http://cite-architecture.github.io
// For a note on CTS citations see https://github.com/cite-architecture/ctsurn_spec/blob/master/md/specification.md
type CTSURN struct {
	ID, Base, Protocol, Namespace, Work, Passage string
	InValid                                      bool
}

// Textgroup is a collection of works
type Textgroup struct {
	TextgroupID string
	Works       []Work
}

// Work is a container for CTS passages that belong to the same work
type Work struct {
	WorkID   string
	Passages []Passage
	Ordered  bool
}

// Passage is the smallest CTSNode
type Passage struct {
	PassageID               string
	Range                   bool
	Text                    EncText
	Index                   int
	First, Last, Prev, Next PassLoc
}

// PassLoc is a container for the ID and
// the Index of a Passage for First, Last, Prev, Next
type PassLoc struct {
	Exists    bool
	PassageID string
	Index     int
}

// EncText is a container for different encodings of the same textual information
type EncText struct {
	TXT, MarkDown, CEX, CSV, XML string
}

// SplitCTS splits a CTS URN in its stem and the passage reference
func SplitCTS(s string) CTSURN {
	if !IsCTSURN(s) {
		return CTSURN{ID: s, InValid: true}
	}
	comps := strings.Split(s, ":")
	return CTSURN{ID: s,
		Base:      comps[0],
		Protocol:  comps[1],
		Namespace: comps[2],
		Work:      comps[3],
		Passage:   comps[4],
		InValid:   false}
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

// IsCTSURN tests whether a string is a valid CTSURN
func IsCTSURN(s string) bool {
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

// IsTextgroupID tests whether the CTSURN points to the textgroup level
func IsTextgroupID(s string) bool {
	if !IsCTSURN(s) {
		return false
	}
	if len(strings.Split(SplitCTS(s).Work, ".")) != 1 {
		return false
	}
	return true
}

// IsWorkID tests whether the CTSURN points to the textgroup level
func IsWorkID(s string) bool {
	if !IsCTSURN(s) {
		return false
	}
	if len(strings.Split(SplitCTS(s).Work, ".")) != 2 {
		return false
	}
	return true
}

// IsVersionID tests whether the CTSURN points to the textgroup level
func IsVersionID(s string) bool {
	if !IsCTSURN(s) {
		return false
	}
	if len(strings.Split(SplitCTS(s).Work, ".")) != 3 {
		return false
	}
	return true
}

// IsExemplarID tests whether the CTSURN points to the textgroup level
func IsExemplarID(s string) bool {
	if !IsCTSURN(s) {
		return false
	}
	if len(strings.Split(SplitCTS(s).Work, ".")) != 4 {
		return false
	}
	return true
}

// GetPassageByID searches for an ID in a given work
func GetPassageByID(id string, w Work) Passage {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return w.Passages[i]
		}
	}
	return Passage{}
}

// GetIndexByID searches for an ID in a given work and returns its Index
func GetIndexByID(id string, w Work) int {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return i
		}
	}
	return 0
}

// GetPassageByInd returns Passage, given an Index and a Work
func GetPassageByInd(i int, w Work) Passage {
	return w.Passages[i]
}

// GetLast returns the last Passage given a Work
func GetLast(w Work) Passage {
	return w.Passages[len(w.Passages)-1]
}

// GetFirst returns the last Passage given a Work
func GetFirst(w Work) Passage {
	return w.Passages[0]
}

// GetNext returns the next Passage given a Work and a PassageID
func GetNext(id string, w Work) Passage {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return w.Passages[w.Passages[i].Prev.Index]
		}
	}
	return Passage{}
}

// GetPrev returns the previous Passage given a Work and a PassageID
func GetPrev(id string, w Work) Passage {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return w.Passages[w.Passages[i].Prev.Index]
		}
	}
	return Passage{}
}

// DelPassage deletes a Passage from a work by changing the references
func DelPassage(id string, w Work) Work {
	index := GetIndexByID(id, w)
	passage := GetPassageByInd(index, w)
	prevInd := passage.Prev.Index
	nextInd := passage.Next.Index
	w.Passages[prevInd].Next = passage.Next
	w.Passages[nextInd].Prev = passage.Prev
	w.Passages[index] = Passage{}
	w.Ordered = false
	return w
}
