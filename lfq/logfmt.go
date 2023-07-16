// Copyright 2023 Matthew Continisio
package lfq

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/kr/logfmt"
)

// Read logfmt input
type LogfmtReader struct {
	l Line
}

func NewLogfmtReader() *LogfmtReader {
	return &LogfmtReader{}
}

func (r *LogfmtReader) Read(b []byte) (Line, error) {
	r.l = NewLine()
	err := logfmt.Unmarshal(b, r)
	return r.l, err
}

func (r *LogfmtReader) HandleLogfmt(key, val []byte) error {
	r.l.m.Set(string(key), string(val))
	return nil
}

// Write logfmt output
type LogfmtWriter struct {
	W io.Writer

	ForceQuote       bool
	DisableQuote     bool
	QuoteEmptyFields bool
}

func NewLogfmtWriter(opts LogfmtWriter) *LogfmtWriter {
	return &LogfmtWriter{os.Stdout, opts.ForceQuote, opts.DisableQuote, opts.QuoteEmptyFields}
}

func (w *LogfmtWriter) Write(l Line) error {
	green := color.New(color.FgGreen)

	b := bytes.Buffer{}
	for pair := l.m.Oldest(); pair != nil; pair = pair.Next() {
		k := pair.Key
		vAny := pair.Value
		v := vAny.(string)

		if b.Len() > 0 {
			b.WriteByte(' ')
		}

		green.Fprint(&b, k)
		b.WriteByte('=')

		if w.needsQuoting(v) {
			b.WriteString(fmt.Sprintf("%q", v))
		} else {
			b.WriteString(v)
		}
	}

	b.WriteByte('\n')

	_, err := w.W.Write(b.Bytes())
	return err
}

func (w *LogfmtWriter) needsQuoting(v string) bool {
	if w.ForceQuote {
		return true
	}
	if w.QuoteEmptyFields && len(v) == 0 {
		return true
	}
	if w.DisableQuote {
		return false
	}

	for _, ch := range v {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}

	return false
}

// Write value-only output
type ValueWriter struct {
	LogfmtWriter
}

func NewValueWriter(opts LogfmtWriter) *ValueWriter {
	return &ValueWriter{LogfmtWriter: LogfmtWriter{os.Stdout, opts.ForceQuote, opts.DisableQuote, opts.QuoteEmptyFields}}
}

func (w *ValueWriter) Write(l Line) error {
	b := bytes.Buffer{}
	for pair := l.m.Oldest(); pair != nil; pair = pair.Next() {
		vAny := pair.Value
		v := vAny.(string)

		if b.Len() > 0 {
			b.WriteByte(' ')
		}

		if w.needsQuoting(v) {
			b.WriteString(fmt.Sprintf("%q", v))
		} else {
			b.WriteString(v)
		}
	}

	b.WriteByte('\n')

	_, err := w.W.Write(b.Bytes())
	return err
}
