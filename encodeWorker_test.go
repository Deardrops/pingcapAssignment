package main

import (
	"os"
	"testing"
)

func TestNewEncodeWorker(t *testing.T) {
	os.Create("input_test.txt")
	NewEncodeWorker(1)
	if !isFileExist("tmp-0") {
		t.Errorf("Failed to create a temp file")
	}
}

func TestEncodeWorker(t *testing.T) {
	wm := WordsMap{}
	wm.Add("bb", 1)
	wm.Add("bc", 2)
	wm.Add("dd", 3)
	encWorker := NewEncodeWorker(2)
	t.Run("save WordsMap", func(t *testing.T) {
		err := encWorker.SaveWordsMap(&wm)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("flush all", func(t *testing.T) {
		err := encWorker.FlushAll()
		if err != nil {
			t.Error(err)
		}
	})
}

func isFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
