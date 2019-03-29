package main

import (
	"os"
	"testing"
)

func TestNewDecodeWorker(t *testing.T) {
	os.Create("tmp-0")
	decWorker := NewDecodeWorker(1)
	decCount := len(decWorker.decoders)
	if decCount != 1 {
		t.Errorf("wrong decoder count, got %d, want %d.\n", decCount, 1)
	}
}

func TestDecodeWorker(t *testing.T) {
	wm := WordsMap{}
	wm.Add("bb", 1)
	wm.Add("bc", 2)
	wm.Add("dd", 3)
	encWorker := NewEncodeWorker(2)
	must(encWorker.SaveWordsMap(&wm))
	must(encWorker.FlushAll())
	decWorker := NewDecodeWorker(2)
	err := decWorker.CloseAllFiles()
	if err != nil {
		t.Error(err)
	}
}
