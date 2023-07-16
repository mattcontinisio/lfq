// Copyright 2023 Matthew Continisio
package lfq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newLine() Line {
	l := NewLine()
	l.m.Set("1", "one")
	l.m.Set("2", "two")
	l.m.Set("3", "three")
	l.m.Set("4", "four")
	return l
}

func TestFilterProcessor(t *testing.T) {
	l1 := newLine()
	k1 := []string{"1", "2"}
	p1 := FilterProcessor{Keys: k1}
	l1 = p1.Process(l1)
	assert.Equal(t, 2, l1.m.Len())
	one, _ := l1.m.Get("1")
	assert.Equal(t, "one", one)
	two, _ := l1.m.Get("2")
	assert.Equal(t, "two", two)

	// Empty keys
	l2 := newLine()
	k2 := []string{}
	p2 := FilterProcessor{Keys: k2}
	l2 = p2.Process(l2)
	assert.Equal(t, 4, l2.m.Len())
}
