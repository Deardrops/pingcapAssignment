package main

import (
	"log"
	"runtime"
)

var DEBUG bool

func debug(f func()) {
	if DEBUG {
		f()
	}
}

var lastTotalFreed uint64

func printMemStats() {
	mb := uint64(1024 * 1024) // MB
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/mb, m.TotalAlloc/mb, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/mb, m.Sys/mb, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}
