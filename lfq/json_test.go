// Copyright 2023 Matthew Continisio
package lfq

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	r := NewJsonReader()
	l, err := r.Read([]byte("{\"one\": 1, \"2\": \"two\", \"three\": \"3 3 3\"}"))
	assert.NoError(t, err)

	assert.Equal(t, 3, l.m.Len())
	one, _ := l.m.Get("one")
	assert.Equal(t, float64(1), one)
	two, _ := l.m.Get("2")
	assert.Equal(t, "two", two)
	three, _ := l.m.Get("three")
	assert.Equal(t, "3 3 3", three)
}

func TestJsonWriter(t *testing.T) {
	l := NewLine()
	l.m.Set("one", 1)
	l.m.Set("2", "two")
	l.m.Set("three", "3 3 3")

	// Write to stdout
	w := NewJsonWriter()
	err := w.Write(l)
	assert.NoError(t, err)

	// Write to buffer
	var b bytes.Buffer
	w = &JsonWriter{&b}
	err = w.Write(l)
	assert.NoError(t, err)
	assert.Equal(t, "{\"one\":1,\"2\":\"two\",\"three\":\"3 3 3\"}\n", b.String())
}
