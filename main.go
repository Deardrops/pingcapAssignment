package main

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	_ "net/http/pprof"
	"os"
)

var inputFlag = flag.String("input", "input.txt", "path to the input file")
var nSlice = flag.Int("count", 10, "number of the slice files")
var defaultMapLen = flag.Int("maplen", 10000, "default length of Map")

func main() {
	// For debugging, you can set an environment variable before running:
	// GODEBUG="gctrace=1"
	// or uncomment the following line (see debug.go for more detail):
	// DEBUG = true

	flag.Parse()
	fmt.Println("------------------ step 1: split input file -------------------------")

	seqTotal, err := SplitInput(*inputFlag, *nSlice)
	if err != nil {
		log.Fatalf("Failed to split input file. %T:%v\n", err, err)
	}

	fmt.Println("------------ step 2: find the first non-repeating word --------------")

	decWorker := NewDecodeWorker(*nSlice)
	defer decWorker.CloseAllFiles()

	globalFirstWord := WordDict{"", CountIndex{0, seqTotal}}

	// Read each slice file in turn and rebuild hashmap
	for _, decoder := range decWorker.decoders {
		UniqueWordsMap := BuildUniqueWordsMap(decoder)

		// find the word with minimum sequence number in a slice file
		firstWord := UniqueWordsMap.FindMinSeqWord(seqTotal)

		// find the word with minimum sequence number in each slice file
		if firstWord.Seq < globalFirstWord.Seq {
			globalFirstWord = firstWord
		}

		UniqueWordsMap = nil
		debug(printMemStats)
	}

	fmt.Println(globalFirstWord)
}

// SplitInput read each word in the file,
// store it in hashmap,
// and write it into different slice file.
func SplitInput(filename string, nSlice int) (int, error) {
	encWorker := NewEncodeWorker(nSlice)
	defer encWorker.CloseAllFile()
	defer encWorker.FlushAll()

	f, err := os.Open(filename)
	if err != nil {
		return -1, errors.New("Failed to read file")
	}
	defer closeFile(f)
	r := bufio.NewReader(f)

	fi, err := f.Stat()
	totalSize := int(fi.Size())
	sliceSize := totalSize / nSlice
	currentSize := 0

	wordsMap := make(WordsMap, *defaultMapLen)
	byts := make([]byte, 0, 32)
	seq := 0 // sequence number of each word

	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				// the input file has been read out
				break
			} else {
				return -1, err
			}
		}
		if isLetter(b) {
			byts = append(byts, b)
		} else {
			if len(byts) != 0 {
				// Save the word to WordMap
				word := string(byts)
				seq++
				wordsMap.Add(word, seq)
				byts = byts[:0]
			}
		}
		if currentSize > sliceSize {
			// for avoiding memory limit exceeded,
			// save WordsMap to disk and free up memory as needed
			must(encWorker.SaveWordsMap(&wordsMap))
			debug(printMemStats)
			wordsMap = make(WordsMap, *defaultMapLen)
			currentSize = 0
		}
		currentSize++
	}
	must(encWorker.SaveWordsMap(&wordsMap))
	debug(printMemStats)
	return seq, nil
}

// BuildUniqueWordsMap read all data in a slice file,
// merge duplicates and rebuild the WordsMap with unique words
func BuildUniqueWordsMap(dec *gob.Decoder) *WordsMap {
	uniqueWordsMap := make(WordsMap, *defaultMapLen)
	for {
		wd := WordDict{}
		err := dec.Decode(&wd)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Failed to read tmp file.%T:%v/n", err, err)
			}
		}

		if ci, ok := uniqueWordsMap[wd.Word]; ok {
			uniqueWordsMap[wd.Word] = CountIndex{wd.Count + ci.Count, ci.Seq}
		} else {
			uniqueWordsMap[wd.Word] = CountIndex{wd.Count, wd.Seq}
		}
	}
	return &uniqueWordsMap
}
