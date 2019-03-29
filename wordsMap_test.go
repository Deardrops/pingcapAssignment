package main

import (
	"reflect"
	"testing"
)

func TestWordsMap_Add(t *testing.T) {
	t.Run("add a new Key", func(t *testing.T) {
		wm := WordsMap{}
		wm.Add("ovo", 1)
		expect := WordsMap{}
		expect["ovo"] = CountIndex{1, 1}
		if !reflect.DeepEqual(wm, expect) {
			t.Errorf("got %+v, want %+v\n", wm, expect)
		}
	})

	t.Run("add an exist Key", func(t *testing.T) {
		wm := WordsMap{}
		wm["ovo"] = CountIndex{1, 10}
		wm.Add("ovo", 20)
		expect := WordsMap{}
		expect["ovo"] = CountIndex{2, 10}
		if !reflect.DeepEqual(wm, expect) {
			t.Errorf("got %+v, want %+v\n", wm, expect)
		}
	})
}

func TestWordsMap_FindMinSeqWord(t *testing.T) {
	wm := WordsMap{}
	wm.Add("eins", 1)
	wm.Add("zwei", 2)
	wm.Add("drei", 3)
	wm.Add("drei", 4)
	wm.Add("eins", 5)
	got := wm.FindMinSeqWord(5)
	expect := WordDict{
		"zwei",
		CountIndex{1, 2},
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %+v, want %+v\n", got, expect)
	}
}
