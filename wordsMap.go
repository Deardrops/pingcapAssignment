package main

import "strings"

// WordDict is the wrapper of word,
// containing the word string, total number of occurrences and sequence number
type WordDict struct {
	Word string
	CountIndex
}

// CountIndex is a wrapper of two attribute of word,
// total number of occurrences and sequence number
type CountIndex struct {
	Count int // total number of occurrences
	Seq   int // sequence number of word
}

// WordsMap is a wrapper of Map, each item is a word
// Key: string of word
// Value: its total number of occurrences and sequence number
type WordsMap map[string]CountIndex

// Add a/an new/exist word to WordsMap
func (wm *WordsMap) Add(word string, seq int) {
	word = strings.ToLower(word)
	if ci, ok := (*wm)[word]; ok {
		(*wm)[word] = CountIndex{ci.Count + 1, ci.Seq}
	} else {
		(*wm)[word] = CountIndex{1, seq}
	}
}

// FindMinSeqWord return the word with minimum sequence number
func (wm *WordsMap) FindMinSeqWord(seqTotal int) WordDict {
	firstWord := WordDict{"", CountIndex{0, seqTotal}}
	for word, ci := range *wm {
		if ci.Count == 1 && ci.Seq < firstWord.Seq {
			firstWord.Word = word
			firstWord.Count = 1
			firstWord.Seq = ci.Seq
		}
	}
	return firstWord
}
