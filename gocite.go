package gocite

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// CiteVerb implemented similar to https://github.com/cite-architecture/cite-verbs/blob/master/cite-collection/cite-verbs-data-draft.csv
type CiteVerb struct {
	ID, Summary, Subject, Object, InverseID string
}

// Cite2Urn implemented as outlined here http://cite-architecture.github.io/2017/02/02/cite2urn_update/
type Cite2Urn struct {
	ID, Base, Protocol, Namespace, Collection, Object string
	InValid                                           bool
}

// CTSURN references text as explained in http://cite-architecture.github.io
// For a note on CTS citations see https://github.com/cite-architecture/ctsurn_spec/blob/master/md/specification.md
type CTSURN struct {
	ID, Base, Protocol, Namespace, Work, Passage string
	InValid                                      bool
}

// Triple is a Simple LinkedData-Triple implementation
type Triple struct {
	Subject, Verb, Object string
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
	ImageLinks              []Triple
}

// PassLoc is a container for the ID and
// the Index of a Passage for First, Last, Prev, Next
type PassLoc struct {
	Exists    bool
	PassageID string
	Index     int
}

// TextAndID is a container for a text-extract and its ID
type TextAndID struct {
	ID, Text string
}

// EncText is a container for different encodings of the same textual information
type EncText struct {
	TXT, Brucheion, MarkDown, CEX, XML, Diplomatic, Normalised string
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

// SplitCITE splits a Cite URN in its stem and the passage reference
func SplitCITE(s string) Cite2Urn {
	if !IsCITEURN(s) {
		return Cite2Urn{ID: s, InValid: true}
	}
	comps := strings.Split(s, ":")
	return Cite2Urn{ID: s,
		Base:       comps[0],
		Protocol:   comps[1],
		Namespace:  comps[2],
		Collection: comps[3],
		Object:     comps[4],
		InValid:    false}
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

// WantSubstr tests whether a the passage part of a URN refers to a substring
func WantSubstr(s string) bool {
	return strings.Contains(s, "@")
}

// IsCTSURN tests whether a string is a valid CTSURN
func IsCTSURN(s string) bool {
	test := strings.Split(s, ":")
	switch {
	case len(test) != 5:
		return false
	case test[0] != "urn":
		return false
	case test[1] != "cts":
		return false
	default:
		return true
	}
}

// IsCITEURN tests whether a string is a valid CITE URN
func IsCITEURN(s string) bool {
	test := strings.Split(s, ":")
	switch {
	case len(test) != 5:
		return false
	case test[0] != "urn":
		return false
	case test[1] != "cite2":
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
func GetPassageByID(id string, w Work) (Passage, error) {
	for i := range w.Passages {
		if w.Passages[i].PassageID == id {
			return w.Passages[i], nil
		}
	}
	return Passage{}, errors.New("couldn't find passage")
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

// RReturnSubStr returns the substring identified by the reverse of @substr[n]. [n] is optional.
func RReturnSubStr(cmd, s string) (string, error) {
	newstr := ""
	switch strings.Contains(cmd, "[") {
	case false:
		add, err := before(s, cmd)
		if err != nil {
			return "", err
		}
		newstr = add + cmd
		return newstr, nil
	case true:
		cmdSl := strings.Split(cmd, "[")
		if len(cmdSl) != 2 {
			return s, errors.New("argument error")
		}
		cmdSl[1] = strings.Replace(cmdSl[1], "]", "", 1)
		n, err := strconv.Atoi(cmdSl[1])
		if err != nil {
			return s, err
		}
		if strings.Count(s, cmdSl[0]) < n {
			return s, errors.New("argument error")
		}
		strSl := strings.SplitN(s, cmdSl[0], n+1)
		strSl = strSl[0:n]
		newstr = strings.Join(strSl, cmdSl[0]) + cmdSl[0]
	}
	return newstr, nil
}

// ReturnSubStr returns the substring identified by @substr[n]. [n] is optional.
func ReturnSubStr(cmd, s string) (string, error) {
	newstr := ""
	switch strings.Contains(cmd, "[") {
	case false:
		add, err := after(s, cmd)
		if err != nil {
			return "", err
		}
		newstr = cmd + add
		return newstr, nil
	case true:
		cmdSl := strings.Split(cmd, "[")
		if len(cmdSl) != 2 {
			return s, errors.New("argument error")
		}
		cmdSl[1] = strings.Replace(cmdSl[1], "]", "", 1)
		n, err := strconv.Atoi(cmdSl[1])
		if err != nil {
			return s, err
		}
		if strings.Count(s, cmdSl[0]) < n {
			return s, errors.New("argument error")
		}
		strSl := strings.SplitN(s, cmdSl[0], n+1)
		newstr = cmdSl[0] + strSl[n]
	}
	return newstr, nil
}

func after(value string, a string) (string, error) {
	// Get substring after a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return "", errors.New(a + " not found in " + value)
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return "", nil
	}
	return value[adjustedPos:], nil
}

func before(value string, a string) (string, error) {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return "", errors.New(a + "not found")
	}
	return value[0:pos], nil
}

// ExtractTextByID extracts the textual information from a Passage or multiple Passages in a Work
func ExtractTextByID(id string, w Work) ([]TextAndID, error) {
	text := []string{}
	extrID := []string{}
	startsub := false
	endsub := false
	startcmd := ""
	endcmd := ""
	if !IsCTSURN(id) {
		return []TextAndID{}, errors.New("urn is not a valid cts urn")
	}
	switch IsRange(id) {
	case false:
		switch WantSubstr(id) {
		case false:
			p, err := GetPassageByID(id, w)
			if err != nil {
				return []TextAndID{}, err
			}
			return []TextAndID{{ID: id, Text: p.Text.TXT}}, nil
		case true:
			idSl := strings.Split(id, "@")
			if len(idSl) != 2 {
				return []TextAndID{}, errors.New("two many @")
			}
			p, err := GetPassageByID(idSl[0], w)
			if err != nil {
				return []TextAndID{}, err
			}
			txt := p.Text.TXT
			_, err = ReturnSubStr(idSl[1], txt)
			if err != nil {
				return []TextAndID{}, err
			}
			reg, err := regexp.Compile(`\[\d\]`)
			if err != nil {
				return []TextAndID{}, err
			}
			idSl[1] = reg.ReplaceAllString(idSl[1], "")
			return []TextAndID{{ID: id, Text: idSl[1]}}, nil
		}
	case true:
		start, end, err := findStartEnd(id)
		firstid := start
		lastid := end
		startRoot := strings.Split(start, "@")
		endRoot := strings.Split(end, "@")
		if startRoot[0] == endRoot[0] {
			if !WantSubstr(start) || !WantSubstr(end) {
				return []TextAndID{}, errors.New("substringing in the same line has the format 1@start-1@end")
			}
			p, err := GetPassageByID(startRoot[0], w)
			if err != nil {
				return []TextAndID{}, err
			}
			p.Text.TXT, err = ReturnSubStr(startRoot[1], p.Text.TXT)
			if err != nil {
				return []TextAndID{}, err
			}
			p.Text.TXT, err = RReturnSubStr(endRoot[1], p.Text.TXT)
			if err != nil {
				return []TextAndID{}, err
			}
			return []TextAndID{{ID: id, Text: p.Text.TXT}}, nil
		}
		if err != nil {
			return []TextAndID{}, err
		}
		if WantSubstr(start) {
			idSl := strings.Split(start, "@")
			if len(idSl) != 2 {
				return []TextAndID{}, errors.New("two many @")
			}
			startsub = true
			start = idSl[0]
			startcmd = idSl[1]
		}
		if WantSubstr(end) {
			idSl2 := strings.Split(end, "@")
			if len(idSl2) != 2 {
				return []TextAndID{}, errors.New("two many @")
			}
			endsub = true
			end = idSl2[0]
			endcmd = idSl2[1]
		}
		switch w.Ordered {
		case true:
			startindex, found := GetIndexByID(start, w)
			endindex, found2 := GetIndexByID(end, w)
			if !found || !found2 {
				return []TextAndID{}, errors.New("passage not found")
			}
			for i := startindex; i < endindex+1; i++ {
				switch i {
				case startindex:
					extrID = append(extrID, firstid)
				case endindex:
					extrID = append(extrID, lastid)
				default:
					extrID = append(extrID, w.Passages[i].PassageID)
				}
				text = append(text, w.Passages[i].Text.TXT)
			}
		case false:
			_, found := GetIndexByID(start, w)
			_, found2 := GetIndexByID(end, w)
			if !found || !found2 {
				return []TextAndID{}, errors.New("passage not found")
			}
			found = false
			startID := start
			IDsVisited := []string{}
			for !found {
				p, err := GetPassageByID(startID, w)
				switch startID {
				case start:
					extrID = append(extrID, firstid)
				case end:
					extrID = append(extrID, lastid)
				default:
					extrID = append(extrID, p.PassageID)
				}
				if err != nil {
					return []TextAndID{}, errors.New("passage not found")
				}
				startID = p.Next.PassageID
				if contains(IDsVisited, startID) {
					return []TextAndID{}, errors.New("work is loopy")
				}
				IDsVisited = append(IDsVisited, startID)
				if p.PassageID == end {
					found = true
				}
				if p.PassageID != end && p.Next.Exists != true {
					return []TextAndID{}, errors.New("unexpected end of work")
				}
				text = append(text, p.Text.TXT)
			}
		}
		if startsub {
			text[0], err = ReturnSubStr(startcmd, text[0])
			if err != nil {
				return []TextAndID{}, err
			}
		}
		if endsub {
			text[len(text)-1], err = RReturnSubStr(endcmd, text[len(text)-1])
			if err != nil {
				return []TextAndID{}, err
			}
		}
	}
	selection := []TextAndID{}
	for i := range text {
		selection = append(selection, TextAndID{ID: extrID[i], Text: text[i]})
	}
	return selection, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findStartEnd(s string) (start, end string, err error) {
	urn := SplitCTS(s)
	if urn.InValid {
		return "", "", errors.New("invalid urn")
	}
	passIDs := strings.Split(urn.Passage, "-")
	if len(passIDs) != 2 {
		return "", "", errors.New("invalid urn")
	}
	start = strings.Join([]string{urn.Base, urn.Protocol, urn.Namespace, urn.Work, passIDs[0]}, ":")
	end = strings.Join([]string{urn.Base, urn.Protocol, urn.Namespace, urn.Work, passIDs[1]}, ":")
	err = nil
	return
}
