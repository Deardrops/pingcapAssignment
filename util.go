package main

import (
	"hash/fnv"
	"log"
	"os"
)

// Check if a byte is legal letter, which means in [a-zA-Z]
func isLetter(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

// An hash function copied from lab1 (MapReduce) of MIT 6.824
func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalf("Failed to close file. %T:%v\n", err, err)
	}
}
