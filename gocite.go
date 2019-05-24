package gocite

import (
	"errors"
	"log"
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
	WorkID      string
	Passages    []Passage
	Ordered     bool
	First, Last PassLoc
}

// Passage is the smallest CTSNode
type Passage struct {
	PassageID  string
	Range      bool
	Text       EncText
	Index      int
	Prev, Next PassLoc
	ImageLinks []Triple
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
func GetPassageByInd(sliceIndex int, work Work) (Passage, error) {
	if sliceIndex >= 0 && sliceIndex <= len(work.Passages)-1 {
		return work.Passages[sliceIndex], nil
	}
	return Passage{}, errors.New("GetPassageByInd: Index out of bounds of work.Passages slice")
}

// GetFirst returns the Passage that is the first in the passage slice in a given a Work
//(not the one saved as Work.First)
func GetFirst(work Work) Passage {
	return work.Passages[0]
}

// GetLast returns the Passage with the last index in the passage slice in a given a Work
//(not the one saved as Work.Last)
func GetLast(work Work) Passage {
	return work.Passages[len(work.Passages)-1]
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
func DelPassage(passageID string, work Work) (Work, error) {
	if len(work.Passages) == 0 {
		return work, errors.New("DelPassage: Work was empty")
	}
	index, found := GetIndexByID(passageID, work)
	if !found {
		return work, errors.New("DelPassage: Passage " + passageID + " not found in Work " + work.WorkID)
	}
	passage, err := GetPassageByInd(index, work)
	if err != nil {
		log.Printf("DelPassage: Error getting Passage by id: %s\n", err)
		return work, err
	}
	switch {
	//passage without Prev and Next: was the last passage remaining in this Work
	case !work.Passages[index].Prev.Exists && !work.Passages[index].Next.Exists:
		temp := Work{WorkID: work.WorkID, Ordered: true} //an empty work that has nothing but a WorkID and Ordered field
		return temp, nil
		//passage has no Prev: was the first Passage in the Work
	case !work.Passages[index].Prev.Exists:
		work, err := DelFirstPassage(work)
		if err == nil {
			return work, nil
		}
		return work, err
		//passage has no Next: was the last Passage in the Work
	case !work.Passages[index].Next.Exists:
		work, err := DelLastPassage(work)
		if err == nil {
			return work, nil
		}
		return work, err
	default:
		work.Passages[passage.Prev.Index].Next = passage.Next
		work.Passages[passage.Next.Index].Prev = passage.Prev
		work.Passages[index] = Passage{}
		work.Ordered = false
	}
	return work, nil
}

// DelFirstPassage deletes the first Passage from a Work by changing the references
func DelFirstPassage(work Work) (Work, error) {
	if len(work.Passages) == 0 {
		return work, errors.New("DelFirstPassage: Work was empty")
	}
	passageIndex, found := GetFirstIndex(work)
	if !found {
		return work, errors.New("DelFirstPassage: First Index not found")
	}
	work.First = work.Passages[passageIndex].Next
	work.Passages[work.First.Index].Prev = PassLoc{Exists: false, PassageID: "", Index: 0}
	work.Passages[passageIndex] = Passage{}
	return work, nil
}

// DelLastPassage deletes the last Passage from a Work by changing the references
func DelLastPassage(work Work) (Work, error) {
	if len(work.Passages) == 0 {
		return work, errors.New("DelLastPassage: Work was empty")
	}
	passageIndex, found := GetLastIndex(work)
	if !found {
		return work, errors.New("DelLastPassage: Last Index not found")
	}
	work.Last = work.Passages[passageIndex].Prev
	work.Passages[work.Last.Index].Next = PassLoc{}
	work.Passages[passageIndex] = Passage{}
	return work, nil
}

/*//FindFirstIndex is deprecated and replaced by GetFirstIndex for legacy
func FindFirstIndex(work Work) (int, bool) {
	return GetFirstIndex(work)
}*/

// FindFirstIndex was returning the First.Index saved in a passage.
//This task is now fulfilled by GetFirstIndex
//Now, FindFirstIndex returns the lowest Passage.Index of the Passages in a Work,
//along with a bool whether it has found one.
//This is necessary for the first analysis of a work. For example in SortPassages.
func FindFirstIndex(work Work) (int, bool) {
	for i := range work.Passages {
		if work.Passages[i].Prev.Exists == false && work.Passages[i].Next.Exists == true {
			return work.Passages[i].Index, true
		}
	}
	return 0, false
}

// GetFirstIndex returns the Work.First.Index of given Work
//along with a bool whether it has found that index.
func GetFirstIndex(work Work) (int, bool) {
	if work.First.Exists {
		return work.First.Index, true
	}
	return 0, false
}

/*//FindLastIndex is deprecated and repalced by GetlastIndex for legacy
func FindLastIndex(work Work) (int, bool) {
	return GetLastIndex(work)
}*/

// FindLastIndex was returning the Last.Index saved in a passage.
//This task is now fulfilled by GetLastIndex
//Now, FindLastIndex returns the highest Passage.Index of the Passages in a Work
//along with a bool whether it has found one.
//This is necessary for the first analysis of a work. For example in SortPassages.
func FindLastIndex(work Work) (int, bool) {
	for i := len(work.Passages) - 1; i >= 0; i-- {
		if work.Passages[i].Prev.Exists == true && work.Passages[i].Next.Exists == false {
			return work.Passages[i].Index, true
		}
	}
	return 0, false
}

// GetLastIndex returns the Work.Last.Index of given Work
//along with a bool whether it has found that index.
func GetLastIndex(work Work) (int, bool) {
	if work.Last.Exists {
		return work.Last.Index, true
	}
	return 0, false
}

// SortPassages sorts the Passages in the Work.Passages slice from First to Last
//according to their Passage.Index values
//empty Passages are not taken over
func SortPassages(work Work) (Work, error) {
	log.Printf("\nSorting Passages in %s\n", work.WorkID)
	if len(work.Passages) == 0 {
		log.Println("Work was empty")
		return work, errors.New("SortPassages: Work was empty")
	}
	if work.Ordered {
		log.Println("Work is marked as ordered")
	} else {
		log.Println("Work is marked as NOT ordered")
	}
	var cursor int
	cursor, found := GetFirstIndex(work) //cursor points to slice index of the next Passage in work which is to be appended to result
	if !found {                          //if there is not first index saved in any of the Passage.Index fields, work will be returned unsorted
		log.Println("First Index not saved in Work yet. Trying to find it now")
		cursor, found = FindFirstIndex(work)
		if !found {
			log.Println("First Index not found in Work")
			return work, errors.New("SortPassages: First Index not found in Work " + work.WorkID) //GetFirstIndex does not actively search for the lowest (and therefore first) index in a work
		}
	}
	log.Printf("First Index: %d, %s\n", cursor, work.Passages[cursor].PassageID)
	lastIndex, found := GetLastIndex(work) //find the last Index, needed as termination condition
	if !found {
		log.Println("Last Index not saved in Work yet. Trying to find it now")
		lastIndex, found = FindLastIndex(work)
		if !found {
			log.Println("Last Index not found in Work")
			return work, errors.New("SortPassages: Last Index not found in work " + work.WorkID)
		}
	}
	log.Printf("Last Index: %d, %s\n", lastIndex, work.Passages[lastIndex].PassageID) //work.Passages[lastIndex].PassageID) shows the Passage at the last index of the USORTED ARRAY!
	result := Work{WorkID: work.WorkID, Ordered: true}                                //result is the sorted Work that will be returned
	index := 0                                                                        //index represents the index the next Passage will get in its Index field
	last := false
	for !last {
		tempPassage := work.Passages[cursor] //get the Passage from work
		log.Printf("\ntempPassage: %s\nIndex: %d\nPrev.Index: %d Prev.PassageID: %s\nNext.Index: %d Next.PassageID: %s\n",
			tempPassage.PassageID, tempPassage.Index,
			tempPassage.Prev.Index, tempPassage.Prev.PassageID,
			tempPassage.Next.Index, tempPassage.Next.PassageID)
		tempPassage.Index = index //set its own Index to the current index
		result.First.Index = 0    //set the index of First to 0
		if index != 0 {           //if this is not the first Passage in result
			tempPassage.Prev.Index = index - 1 //set the Prev index field to one lower than the current index
		} else {
			log.Println("*** I'm the first. ***")
		}
		if index == lastIndex { //if the cursor points to the last index
			last = true //mark it in the last variable
			log.Println("*** I'm the last.  ***")
		}
		if last == false { //if this is not the last Passage in the work
			tempPassage.Next.Index = index + 1                                        //set the Next index field to one higher than the current index
			nextCursor, _ := GetIndexByID(work.Passages[cursor].Next.PassageID, work) //get the Index of the next Passage according to what is saved in the Passage.PassageID the cursor currently points to
			cursor = nextCursor                                                       //set the cursor to the index of the next Passage
			index++                                                                   //increment the index variable
		}
		result.Passages = append(result.Passages, tempPassage) //append the temporary Passage to the resulting work
	}
	result.First = PassLoc{Exists: true, PassageID: result.Passages[0].PassageID, Index: 0}
	result.Passages[0].Prev = PassLoc{}
	result.Last = PassLoc{Exists: true, PassageID: result.Passages[lastIndex].PassageID, Index: lastIndex}
	result.Passages[lastIndex].Next = PassLoc{}
	log.Println("*** Sorting finished. ***")
	log.Println("** Before **")
	log.Printf("work.First.Index: %d, work.First.PassageID: %s\n", work.First.Index, work.First.PassageID)
	log.Printf("work.Last.Index: %d, work.Last.PassageID: %s\n", work.Last.Index, work.Last.PassageID)
	log.Println("** Afterwards **")
	log.Printf("result.First.Index: %d, result.First.PassageID: %s\n", result.First.Index, result.First.PassageID)
	log.Printf("result.Last.Index: %d, result.Last.PassageID: %s\n", result.Last.Index, result.Last.PassageID)
	return result, nil
}

// InsertPassage inserts a Passage into a Work
func InsertPassage(passage Passage, work Work) (Work, error) {
	if len(work.Passages) == 0 { //if the work has no passages yet
		work.First = PassLoc{Exists: true, PassageID: passage.PassageID, Index: 0}
		work.Last = PassLoc{Exists: true, PassageID: passage.PassageID, Index: 0}
		passage.Next = PassLoc{}
		passage.Prev = PassLoc{}
		work.Passages = append(work.Passages, passage)
		return work, nil
	}
	nextIndex, nextExists := GetIndexByID(passage.Next.PassageID, work)
	prevIndex, prevExists := GetIndexByID(passage.Prev.PassageID, work)
	firstIndex, found := GetFirstIndex(work)
	if !found { //we could add a FindFirstIndex here.
		return work, errors.New("InsertPassage: First Index not found")
	}
	lastIndex, found := GetLastIndex(work)
	if !found { //we could add a FindLastIndex here
		return work, errors.New("InsertPassage: Last Index not found")
	}
	passloc := PassLoc{Exists: true, PassageID: passage.PassageID, Index: len(work.Passages)}
	work.First = PassLoc{Exists: true, PassageID: work.Passages[firstIndex].PassageID, Index: firstIndex}
	work.Last = PassLoc{Exists: true, PassageID: work.Passages[lastIndex].PassageID, Index: lastIndex}
	switch {
	case prevExists && !nextExists: //if the new passage is the (new) last passage
		passage.Prev = PassLoc{Exists: true, PassageID: work.Passages[prevIndex].PassageID, Index: prevIndex}
		passage.Next = PassLoc{}
		work.Passages = append(work.Passages, passage)
		work.Passages[prevIndex].Next = passloc
		work.Last = passloc
	case !prevExists && nextExists: //if the new passage is the (new) first passage
		passage.Prev = PassLoc{}
		passage.Next = PassLoc{Exists: true, PassageID: work.Passages[nextIndex].PassageID, Index: nextIndex}
		work.Passages = append(work.Passages, passage)
		work.Passages[nextIndex].Prev = passloc
		work.First = passloc
	default: //if new passage is to be added somewhere in between.
		passage.Prev = PassLoc{Exists: true, PassageID: work.Passages[prevIndex].PassageID, Index: prevIndex}
		passage.Next = PassLoc{Exists: true, PassageID: work.Passages[nextIndex].PassageID, Index: nextIndex}
		work.Passages = append(work.Passages, passage)
		work.Passages[prevIndex].Next = passloc
		work.Passages[nextIndex].Prev = passloc
		work.Passages[len(work.Passages)-1].Index = len(work.Passages) - 1
	}
	work.Ordered = false
	return work, nil
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

// before returns the substring of the originalString that antecedes the beforThisString,
//as long as the beforeThisString is contained in the originalString
func before(originalString string, beforeThisString string) (string, error) {
	// Get substring before a string.
	pos := strings.Index(originalString, beforeThisString)
	if pos == -1 {
		return "", errors.New(beforeThisString + "not found")
	}
	return originalString[0:pos], nil
}

// after returns the substring of the originalString that precedes the afterThisString,
//as long as the afterThisString is contained in the originalString.
func after(originalString string, afterThisString string) (string, error) {
	// Get substring after a string.
	pos := strings.Index(originalString, afterThisString)
	if pos == -1 {
		return "", errors.New(afterThisString + " not found in " + originalString)
	}
	adjustedPos := pos + len(afterThisString)
	if adjustedPos >= len(originalString) {
		return "", nil
	}
	return originalString[adjustedPos:], nil
}

// ExtractTextByID extracts the textual information from a Passage or multiple Passages in a Work
func ExtractTextByID(ctsID string, work Work) ([]TextAndID, error) {
	text := []string{}
	extrID := []string{}
	startsub := false
	endsub := false
	startcmd := ""
	endcmd := ""
	if !IsCTSURN(ctsID) {
		return []TextAndID{}, errors.New("urn is not a valid cts urn")
	}
	switch IsRange(ctsID) {
	case false:
		switch WantSubstr(ctsID) {
		case false:
			p, err := GetPassageByID(ctsID, work)
			if err != nil {
				return []TextAndID{}, err
			}
			return []TextAndID{{ID: ctsID, Text: p.Text.TXT}}, nil
		case true:
			idSl := strings.Split(ctsID, "@")
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
			return []TextAndID{{ID: ctsID, Text: idSl[1]}}, nil
		}
	case true:
		start, end, err := findStartEnd(ctsID)
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
			return []TextAndID{{ID: ctsID, Text: p.Text.TXT}}, nil
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
			if !found {
				return []TextAndID{}, errors.New("Index of start passage not found")
			}
			if !found2 {
				return []TextAndID{}, errors.New("Index of end passage not found")
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

//contains returns true if the 'needle' string is found in the 'haystack' string slice
func contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
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
