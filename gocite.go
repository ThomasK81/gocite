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
// It also returns a bool whether it has found the Passage
func GetIndexByID(id string, w Work) (int, bool) {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return i, true
		}
	}
	return 0, false
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
	if len(w.Passages) == 0 {
		return w
	}
	index, found := GetIndexByID(id, w)
	if !found {
		return w
	}
	passage := GetPassageByInd(index, w)
	switch {
	case !w.Passages[index].Prev.Exists && !w.Passages[index].Next.Exists:
		temp := Work{WorkID: w.WorkID, Ordered: true}
		return temp
	case !w.Passages[index].Prev.Exists:
		w := DelFirstPassage(w)
		return w
	case !w.Passages[index].Next.Exists:
		w := DelLastPassage(w)
		return w
	default:
		prevInd := passage.Prev.Index
		nextInd := passage.Next.Index
		w.Passages[prevInd].Next = passage.Next
		w.Passages[nextInd].Prev = passage.Prev
		w.Passages[index] = Passage{}
		w.Ordered = false
	}
	return w
}

// DelFirstPassage deletes the first Passage from a work by changing the references
func DelFirstPassage(w Work) Work {
	if len(w.Passages) == 0 {
		return w
	}
	passageIndex, found := FindFirstIndex(w)
	if !found {
		return w
	}
	newFirst := w.Passages[passageIndex].Next
	for i := range w.Passages {
		w.Passages[i].First = newFirst
	}
	w.Passages[newFirst.Index].Prev = PassLoc{}
	w.Passages[passageIndex] = Passage{}
	return w
}

// DelLastPassage deletes the last Passage from a work by changing the references
func DelLastPassage(w Work) Work {
	if len(w.Passages) == 0 {
		return w
	}
	passageIndex, found := FindLastIndex(w)
	if !found {
		return w
	}
	newLast := w.Passages[passageIndex].Prev
	for i := range w.Passages {
		w.Passages[i].Last = newLast
	}
	w.Passages[newLast.Index].Next = PassLoc{}
	w.Passages[passageIndex] = Passage{}
	return w
}

// FindFirstIndex returns the first Index of the Passages in a Work.
// It also returns a bool whether it has found one.
func FindFirstIndex(w Work) (int, bool) {
	for i := range w.Passages {
		if w.Passages[i].First.Exists {
			return w.Passages[i].First.Index, true
		}
	}
	return 0, false
}

// FindLastIndex returns the first Index of the Passages in a Work.
// It also returns a bool whether it has found one.
func FindLastIndex(w Work) (int, bool) {
	for i := range w.Passages {
		if w.Passages[i].Last.Exists {
			return w.Passages[i].Last.Index, true
		}
	}
	return 0, false
}

// SortPassages sorts the Passages in a Work from First to Last, empty Passages are being deleted
func SortPassages(w Work) Work {
	if len(w.Passages) == 0 {
		return w
	}
	cursor, found := FindFirstIndex(w)
	if !found {
		return w
	}
	result := Work{WorkID: w.WorkID, Ordered: true}
	index := 0
	last := false
	for !last {
		temp := w.Passages[cursor]
		temp.Index = index
		temp.First.Index = 0
		if index != 0 {
			temp.Prev.Index = index - 1
		}
		if w.Passages[cursor].PassageID == w.Passages[cursor].Last.PassageID {
			last = true
		}
		if last == false {
			temp.Next.Index = index + 1
			cursor = w.Passages[cursor].Next.Index
			index++
		}
		result.Passages = append(result.Passages, temp)
	}
	for i := range result.Passages {
		result.Passages[i].Last.Index = index
	}
	return result
}

// InsertPassage inserts a Passage into a Work
func InsertPassage(p Passage, w Work) Work {
	if len(w.Passages) == 0 {
		p.First = PassLoc{Exists: true, PassageID: p.PassageID, Index: 0}
		p.Last = PassLoc{Exists: true, PassageID: p.PassageID, Index: 0}
		p.Next = PassLoc{}
		p.Prev = PassLoc{}
		w.Passages = append(w.Passages, p)
		return w
	}
	nextIndex, nextExists := GetIndexByID(p.Next.PassageID, w)
	prevIndex, prevExists := GetIndexByID(p.Prev.PassageID, w)
	firstIndex, _ := FindFirstIndex(w)
	lastIndex, _ := FindLastIndex(w)
	passloc := PassLoc{Exists: true, PassageID: p.PassageID, Index: len(w.Passages)}
	p.First = PassLoc{Exists: true, PassageID: w.Passages[firstIndex].PassageID, Index: firstIndex}
	p.Last = PassLoc{Exists: true, PassageID: w.Passages[lastIndex].PassageID, Index: lastIndex}
	switch {
	case !nextExists && prevExists:
		p.Next = PassLoc{}
		p.Prev = PassLoc{Exists: true, PassageID: w.Passages[prevIndex].PassageID, Index: prevIndex}
		w.Passages = append(w.Passages, p)
		w.Passages[prevIndex].Next = passloc
		for i := range w.Passages {
			w.Passages[i].Last = passloc
		}
	case nextExists && !prevExists:
		p.Prev = PassLoc{}
		p.Next = PassLoc{Exists: true, PassageID: w.Passages[nextIndex].PassageID, Index: nextIndex}
		w.Passages = append(w.Passages, p)
		w.Passages[nextIndex].Prev = passloc
		for i := range w.Passages {
			w.Passages[i].First = passloc
		}
	default:
		p.Next = PassLoc{Exists: true, PassageID: w.Passages[nextIndex].PassageID, Index: nextIndex}
		p.Prev = PassLoc{Exists: true, PassageID: w.Passages[prevIndex].PassageID, Index: prevIndex}
		w.Passages = append(w.Passages, p)
		w.Passages[prevIndex].Next = passloc
		w.Passages[nextIndex].Prev = passloc
		w.Passages[len(w.Passages)-1].Index = len(w.Passages) - 1
	}
	w.Ordered = false
	return w
}
