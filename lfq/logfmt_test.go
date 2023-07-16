// Copyright 2023 Matthew Continisio
package lfq

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogfmtReader(t *testing.T) {
	r := NewLogfmtReader()
	l, err := r.Read([]byte("one=1 2=two three=\"3 3 3\""))
	assert.NoError(t, err)

	assert.Equal(t, 3, l.m.Len())
	one, _ := l.m.Get("one")
	assert.Equal(t, "1", one)
	two, _ := l.m.Get("2")
	assert.Equal(t, "two", two)
	three, _ := l.m.Get("three")
	assert.Equal(t, "3 3 3", three)
}

func TestLogfmtWriter(t *testing.T) {
	l := NewLine()
	l.m.Set("one", "1")
	l.m.Set("2", "two")
	l.m.Set("three", "3 3 3")

	// Write to stdout
	w := NewLogfmtWriter(LogfmtWriter{})
	err := w.Write(l)
	assert.NoError(t, err)

	// Write to buffer
	var b bytes.Buffer
	w = &LogfmtWriter{W: &b}
	err = w.Write(l)
	assert.NoError(t, err)
	assert.Equal(t, "one=1 2=two three=\"3 3 3\"\n", b.String())
}

func TestValueWriter(t *testing.T) {
	l := NewLine()
	l.m.Set("one", "1")
	l.m.Set("2", "two")
	l.m.Set("three", "3 3 3")

	// Write to stdout
	w := NewValueWriter(LogfmtWriter{})
	err := w.Write(l)
	assert.NoError(t, err)

	// Write to buffer
	var b bytes.Buffer
	w = &ValueWriter{LogfmtWriter: LogfmtWriter{W: &b}}
	err = w.Write(l)
	assert.NoError(t, err)
	assert.Equal(t, "1 two \"3 3 3\"\n", b.String())
}
