package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"reflect"
	"testing"
)

// Recursively generates the full arrangement of character arrays
// and outputs them to the file with space division.
func permute(chars []byte, curr []byte, w *bufio.Writer) {
	if len(chars) == 1 {
		curr = append(curr, chars[0])
		w.Write(curr)
		w.WriteByte(' ')
		return
	}

	for i, num := range chars {
		tmp := append([]byte{}, chars[:i]...)
		tmp = append(tmp, chars[i+1:]...)
		permute(tmp, append(curr, num), w)
	}
}

// Outputs the full arrangement of the specified
// character slice into the input file
func createTestInput(chars []byte) {
	f, err := os.Create("input_test.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	permute(chars, []byte{}, w)
	must(w.Flush())
}

func clearTestInput() {
	must(os.Remove("input_test.txt"))
}

func TestSplitInput(t *testing.T) {
	createTestInput([]byte("ABCDEFGHI"))
	defer clearTestInput()
	nSlice := 3
	seqTotal, err := SplitInput("input_test.txt", nSlice)
	if err != nil {
		t.Error(err)
	}
	expectSeqTotal := 362880
	if seqTotal != expectSeqTotal {
		t.Errorf("seq total wrong, got %d, want %d", seqTotal, expectSeqTotal)
	}
}

func TestBuildUniqueWordsMap(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(b)
	enc := gob.NewEncoder(w)
	must(enc.Encode(WordDict{"bb", CountIndex{1, 1}}))
	must(enc.Encode(WordDict{"bc", CountIndex{2, 2}}))
	must(enc.Encode(WordDict{"bb", CountIndex{1, 3}}))
	must(enc.Encode(WordDict{"ca", CountIndex{1, 4}}))
	must(w.Flush())
	dec := gob.NewDecoder(bufio.NewReader(b))
	uniqueWordsMap := BuildUniqueWordsMap(dec)
	expect := WordsMap{}
	expect["bb"] = CountIndex{2, 1}
	expect["ca"] = CountIndex{1, 4}
	expect["bc"] = CountIndex{2, 2}

	if !reflect.DeepEqual(*uniqueWordsMap, expect) {
		t.Errorf("got %+v, want %+v", *uniqueWordsMap, expect)
	}
}
