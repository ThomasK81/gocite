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

// Cite2Urn implemented as outlined here http://cite-architecture.org/cite2urn/
type Cite2Urn struct {
	ID, Base, Protocol, Namespace, Collection, Object string
	InValid                                           bool
}

// CTSURN references text as explained in http://cite-architecture.org/ctsurn/
// For a note on CTS citations see http://cite-architecture.github.io/cts_spec/
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

// SplitCTS splits a CTS URN string in its stem and the passage reference
//and returns it as a CTSURN
func SplitCTS(URNString string) CTSURN {
	if !IsCTSURN(URNString) {
		return CTSURN{ID: URNString, InValid: true}
	}
	comps := strings.Split(URNString, ":")
	return CTSURN{ID: URNString,
		Base:      comps[0],
		Protocol:  comps[1],
		Namespace: comps[2],
		Work:      comps[3],
		Passage:   comps[4],
		InValid:   false}
}

// SplitCITE splits a Cite URN string in its stem and the passage reference
//and returns it as a Cite2Urn
func SplitCITE(URNString string) Cite2Urn {
	if !IsCITEURN(URNString) {
		return Cite2Urn{ID: URNString, InValid: true}
	}
	comps := strings.Split(URNString, ":")
	return Cite2Urn{ID: URNString,
		Base:       comps[0],
		Protocol:   comps[1],
		Namespace:  comps[2],
		Collection: comps[3],
		Object:     comps[4],
		InValid:    false}
}

// IsRange returns a boolean whether a CTS URN string is a range
func IsRange(URNString string) bool {
	switch {
	case len(strings.Split(URNString, ":")) < 5:
		return false
	case strings.Contains(strings.Split(URNString, ":")[4], "-"):
		return true
	default:
		return false
	}
}

// WantSubstr tests whether the passage part of a URN string refers to a substring
func WantSubstr(URNString string) bool {
	return strings.Contains(URNString, "@")
}

// IsCTSURN tests whether a string is a valid CTSURN
func IsCTSURN(URNString string) bool {
	test := strings.Split(URNString, ":")
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
func IsCITEURN(URNString string) bool {
	test := strings.Split(URNString, ":")
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

// IsTextgroupID tests whether a CTSURN (string) points to the textgroup level
func IsTextgroupID(URNString string) bool {
	if !IsCTSURN(URNString) {
		return false
	}
	if len(strings.Split(SplitCTS(URNString).Work, ".")) != 1 {
		return false
	}
	return true
}

// IsWorkID tests whether a CTSURN (string) points to the work level
func IsWorkID(URNString string) bool {
	if !IsCTSURN(URNString) {
		return false
	}
	if len(strings.Split(SplitCTS(URNString).Work, ".")) != 2 {
		return false
	}
	return true
}

// IsVersionID tests whether a CTSURN (string) points to the version level
func IsVersionID(URNString string) bool {
	if !IsCTSURN(URNString) {
		return false
	}
	if len(strings.Split(SplitCTS(URNString).Work, ".")) != 3 {
		return false
	}
	return true
}

// IsExemplarID tests whether a CTSURN (string) points to the exemplar level
func IsExemplarID(URNString string) bool {
	if !IsCTSURN(URNString) {
		return false
	}
	if len(strings.Split(SplitCTS(URNString).Work, ".")) != 4 {
		return false
	}
	return true
}

// GetPassageByID returns a Passage given the PassageID in a given Work
func GetPassageByID(passageID string, work Work) (Passage, error) {
	for i := range work.Passages {
		if work.Passages[i].PassageID == passageID {
			return work.Passages[i], nil
		}
	}
	return Passage{}, errors.New("couldn't find passage")
}

// GetIndexByID searches for an ID in a given work and if found,
// returns its slice index in the Work.Passages slice (not the passage.Index)
// along with a bool indicating whether it has found the Passage
func GetIndexByID(passageID string, work Work) (int, bool) {
	for index := range work.Passages {
		if work.Passages[index].PassageID == passageID {
			return index, true
		}
	}
	return 0, false
}

// GetPassageByInd returns the Passage at given Index in the Work.Passages slice
//(Not by Passage.Index)
func GetPassageByInd(sliceIndex int, work Work) Passage {
	return work.Passages[sliceIndex]
}

// GetLast returns the Passage that appears to be the last in the passage slice in a given a Work
//(not the one saved as Passage.Last)
func GetLast(work Work) Passage {
	return work.Passages[len(work.Passages)-1]
}

// GetFirst returns the Passage that appears to be the first in the passage slice in a given a Work
//(not the one saved as Passage.First)
func GetFirst(work Work) Passage {
	return work.Passages[0]
}

// GetNext returns the Passage after the Passage given the PassageID in given Work
//(not the next in the work.Passages slice)
func GetNext(passageID string, work Work) Passage {
	for i := range work.Passages {
		if work.Passages[i].PassageID == passageID {
			return work.Passages[work.Passages[i].Next.Index]
		}
	}
	return Passage{}
}

// GetPrev returns the Passage previous to the given the PassageID in given a Work
//(not the next in the work.Passages slice)
func GetPrev(passageID string, work Work) Passage {
	for i := range work.Passages {
		if work.Passages[i].PassageID == passageID {
			return work.Passages[work.Passages[i].Prev.Index]
		}
	}
	return Passage{}
}

// DelPassage deletes a Passage from a Work by changing the references
func DelPassage(passageID string, work Work) Work {
	if len(work.Passages) == 0 {
		return work
	}
	index, found := GetIndexByID(passageID, work)
	if !found {
		return work
	}
	passage := GetPassageByInd(index, work)
	switch {
	case !work.Passages[index].Prev.Exists && !work.Passages[index].Next.Exists:
		temp := Work{WorkID: work.WorkID, Ordered: true}
		return temp
	case !work.Passages[index].Prev.Exists:
		work := DelFirstPassage(work)
		return work
	case !work.Passages[index].Next.Exists:
		work := DelLastPassage(work)
		return work
	default:
		prevInd := passage.Prev.Index
		nextInd := passage.Next.Index
		work.Passages[prevInd].Next = passage.Next
		work.Passages[nextInd].Prev = passage.Prev
		work.Passages[index] = Passage{}
		work.Ordered = false
	}
	return work
}

// DelFirstPassage deletes the first Passage from a Work by changing the references
func DelFirstPassage(work Work) Work {
	if len(work.Passages) == 0 {
		return work
	}
	passageIndex, found := GetFirstIndex(work)
	if !found {
		return work
	}
	newFirst := work.Passages[passageIndex].Next
	for i := range work.Passages {
		work.Passages[i].First = newFirst
	}
	work.Passages[newFirst.Index].Prev = PassLoc{}
	work.Passages[passageIndex] = Passage{}
	return work
}

// DelLastPassage deletes the last Passage from a Work by changing the references
func DelLastPassage(work Work) Work {
	if len(work.Passages) == 0 {
		return work
	}
	passageIndex, found := GetLastIndex(work)
	if !found {
		return work
	}
	newLast := work.Passages[passageIndex].Prev
	for i := range work.Passages {
		work.Passages[i].Last = newLast
	}
	work.Passages[newLast.Index].Next = PassLoc{}
	work.Passages[passageIndex] = Passage{}
	return work
}

// GetFirstIndex returns the Passage.First.Index of the first Passage in a Work
//that has Passage.First.Index.Exists set to true,
//along with a bool whether it has found that index.
func GetFirstIndex(work Work) (int, bool) {
	for i := range work.Passages {
		if work.Passages[i].First.Exists {
			return work.Passages[i].First.Index, true
		}
	}
	return 0, false
}

// GetLastIndex returns the Passage.Last.Index of the first Passage in a Work
//that has Passage.Last.Index.Exists set to true,
//along with a bool whether it has found that index.
func GetLastIndex(work Work) (int, bool) {
	for i := range work.Passages {
		if work.Passages[i].Last.Exists {
			return work.Passages[i].Last.Index, true
		}
	}
	return 0, false
}

// SortPassages sorts the Passages in the Work.Passages slice from First to Last
//according to their Passage.Index values
//empty Passages are not being appended
func SortPassages(work Work) Work {
	if len(work.Passages) == 0 {
		return work
	}
	cursor, found := GetFirstIndex(work)
	if !found {
		return work
	}
	result := Work{WorkID: work.WorkID, Ordered: true}
	index := 0
	last := false
	for !last {
		temp := work.Passages[cursor]
		temp.Index = index
		temp.First.Index = 0
		if index != 0 {
			temp.Prev.Index = index - 1
		}
		if work.Passages[cursor].PassageID == work.Passages[cursor].Last.PassageID {
			last = true
		}
		if last == false {
			temp.Next.Index = index + 1
			cursor = work.Passages[cursor].Next.Index
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
func InsertPassage(passage Passage, work Work) Work {
	if len(work.Passages) == 0 { //if the work has no passages yet
		passage.First = PassLoc{Exists: true, PassageID: passage.PassageID, Index: 0}
		passage.Last = PassLoc{Exists: true, PassageID: passage.PassageID, Index: 0}
		passage.Next = PassLoc{}
		passage.Prev = PassLoc{}
		work.Passages = append(work.Passages, passage)
		return work
	}
	nextIndex, nextExists := GetIndexByID(passage.Next.PassageID, work)
	prevIndex, prevExists := GetIndexByID(passage.Prev.PassageID, work)
	firstIndex, _ := GetFirstIndex(work)
	lastIndex, _ := GetLastIndex(work)
	passloc := PassLoc{Exists: true, PassageID: passage.PassageID, Index: len(work.Passages)}
	passage.First = PassLoc{Exists: true, PassageID: work.Passages[firstIndex].PassageID, Index: firstIndex}
	passage.Last = PassLoc{Exists: true, PassageID: work.Passages[lastIndex].PassageID, Index: lastIndex}
	switch {
	case prevExists && !nextExists: //if the new passage is the (new) last passage
		passage.Prev = PassLoc{Exists: true, PassageID: work.Passages[prevIndex].PassageID, Index: prevIndex}
		passage.Next = PassLoc{}
		work.Passages = append(work.Passages, passage)
		work.Passages[prevIndex].Next = passloc
		for i := range work.Passages {
			work.Passages[i].Last = passloc
		}
	case !prevExists && nextExists: //if the new passage is the (new) first passage
		passage.Prev = PassLoc{}
		passage.Next = PassLoc{Exists: true, PassageID: work.Passages[nextIndex].PassageID, Index: nextIndex}
		work.Passages = append(work.Passages, passage)
		work.Passages[nextIndex].Prev = passloc
		for i := range work.Passages {
			work.Passages[i].First = passloc
		}
	default: //if new passage is to be added somewhere in between.
		passage.Prev = PassLoc{Exists: true, PassageID: work.Passages[prevIndex].PassageID, Index: prevIndex}
		passage.Next = PassLoc{Exists: true, PassageID: work.Passages[nextIndex].PassageID, Index: nextIndex}
		work.Passages = append(work.Passages, passage)
		work.Passages[prevIndex].Next = passloc
		work.Passages[nextIndex].Prev = passloc
		work.Passages[len(work.Passages)-1].Index = len(work.Passages) - 1
	}
	work.Ordered = false
	return work
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
func ExtractTextByID(id string, work Work) ([]TextAndID, error) {
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
			p, err := GetPassageByID(id, work)
			if err != nil {
				return []TextAndID{}, err
			}
			return []TextAndID{{ID: id, Text: p.Text.TXT}}, nil
		case true:
			idSl := strings.Split(id, "@")
			if len(idSl) != 2 {
				return []TextAndID{}, errors.New("two many @")
			}
			p, err := GetPassageByID(idSl[0], work)
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
			p, err := GetPassageByID(startRoot[0], work)
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
		switch work.Ordered {
		case true:
			startindex, found := GetIndexByID(start, work)
			endindex, found2 := GetIndexByID(end, work)
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
					extrID = append(extrID, work.Passages[i].PassageID)
				}
				text = append(text, work.Passages[i].Text.TXT)
			}
		case false:
			_, found := GetIndexByID(start, work)
			_, found2 := GetIndexByID(end, work)
			if !found || !found2 {
				return []TextAndID{}, errors.New("passage not found")
			}
			found = false
			startID := start
			IDsVisited := []string{}
			for !found {
				p, err := GetPassageByID(startID, work)
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

//contains returns true if the 'needle' string is found in the 'heystack' string slice
func contains(heystack []string, needle string) bool {
	for _, straw := range heystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func findStartEnd(URNString string) (start, end string, err error) {
	urn := SplitCTS(URNString)
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
