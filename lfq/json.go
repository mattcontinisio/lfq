// Copyright 2023 Matthew Continisio
package lfq

import (
	"encoding/json"
	"io"
	"os"
)

// Read JSON input
type JsonReader struct{}

func NewJsonReader() *JsonReader {
	return &JsonReader{}
}

func (r *JsonReader) Read(b []byte) (Line, error) {
	l := NewLine()
	err := json.Unmarshal(b, &l.m)
	return l, err
}

// Write JSON output
type JsonWriter struct {
	W io.Writer
}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{os.Stdout}
}

func (w *JsonWriter) Write(l Line) error {
	e := json.NewEncoder(w.W)
	return e.Encode(l.m)
}
