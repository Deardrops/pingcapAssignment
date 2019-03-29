package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
	"strconv"
)

type DecodeWorker struct {
	files    []*os.File
	decoders []*gob.Decoder
}

func NewDecodeWorker(nSlice int) *DecodeWorker {
	files := make([]*os.File, nSlice)
	decoders := make([]*gob.Decoder, nSlice)
	for i := 0; i < nSlice; i++ {
		f, err := os.Open("tmp-" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		files[i] = f
		decoders[i] = gob.NewDecoder(bufio.NewReader(f))
	}
	return &DecodeWorker{
		files,
		decoders,
	}
}

func (dw *DecodeWorker) CloseAllFiles() error {
	for _, f := range dw.files {
		err := f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
