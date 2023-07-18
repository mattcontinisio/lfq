// Copyright 2023 Matthew Continisio
package lfq

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// A parsed line
type Line struct {
	m *orderedmap.OrderedMap[string, any]
}

func NewLine() Line {
	return Line{orderedmap.New[string, any]()}
}

// Parses bytes into Lines
type Reader interface {
	Read(b []byte) (Line, error)
}

// Writes lines to stdout
type Writer interface {
	Write(l Line) error
}

// Processes lines in some way (e.g. filter, transform, etc.)
type Processor interface {
	Process(l Line) Line
}

// Filters keys
type FilterProcessor struct {
	Keys []string
}

func (p *FilterProcessor) Process(l Line) Line {
	if len(p.Keys) == 0 {
		return l
	}

	for pair := l.m.Oldest(); pair != nil; {
		k := pair.Key
		pair = pair.Next()
		found := false
		for _, filterKey := range p.Keys {
			if filterKey == k {
				found = true
				break
			}
		}

		if !found {
			l.m.Delete(k)
		}
	}

	return l
}
