package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
)

type EncodeWorker struct {
	files    []*os.File
	writers  []*bufio.Writer
	encoders []*gob.Encoder
	nSlice   int
}

func NewEncodeWorker(nSlice int) *EncodeWorker {
	// Initialize all required resources
	files := make([]*os.File, nSlice)
	writers := make([]*bufio.Writer, nSlice)
	encoders := make([]*gob.Encoder, nSlice)
	for i := 0; i < nSlice; i++ {
		f, err := os.Create("tmp-" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		files[i] = f
		writers[i] = bufio.NewWriter(f)
		encoders[i] = gob.NewEncoder(writers[i])
	}
	return &EncodeWorker{
		files,
		writers,
		encoders,
		nSlice,
	}
}

// SaveWordsMap map and save current wordsMap to different file slices
// through hash function and remainder operation
func (ew *EncodeWorker) SaveWordsMap(wordsMap *WordsMap) error {
	for word, ci := range *wordsMap {
		idx := ihash(word) % ew.nSlice
		wordDict := WordDict{word, ci}
		err := ew.encoders[idx].Encode(wordDict)
		if err != nil {
			return fmt.Errorf("Failed to encode WordDict. %T:%v\n", err, err)
		}
	}
	return nil
}

func (ew *EncodeWorker) FlushAll() error {
	for _, w := range ew.writers {
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (ew *EncodeWorker) CloseAllFile() error {
	for _, f := range ew.files {
		err := f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
