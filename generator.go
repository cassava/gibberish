// Copyright (c) 2015, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package gibberish

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"time"
)

type Generator struct {
	Alphabet *Alphabet

	WordLen      *Normal
	ClauseLen    *Normal
	SentenceLen  *Normal
	ParagraphLen *Normal
	DocumentLen  *Normal

	Prefix string
}

func NewGenerator() *Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Generator{
		Alphabet:     ASCII,
		WordLen:      &Normal{r, 4, 7},
		ClauseLen:    &Normal{r, 5, 4},
		SentenceLen:  &Normal{r, 1.8, 2},
		ParagraphLen: &Normal{r, 3, 2.5},
		DocumentLen:  &Normal{r, 3, 2},
		Prefix:       "# The following text is gibberish, and is for testing purposes.\n#\n\n",
	}
}

func (g *Generator) WriteToFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = g.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) WriteTo(w io.Writer) (int64, error) {
	m, err := g.writeDocument(w)
	return int64(m), err
}

func (g *Generator) writeDocument(w io.Writer) (int, error) {
	return writeRand(g.DocumentLen, 1, func(first, last bool) (n int, err error) {
		var j int
		if first {
			if g.Prefix != "" {
				j, err = io.WriteString(w, g.Prefix)
				n += j
			}
		}
		j, err = g.writeParagraph(w)
		n += j
		if err != nil {
			return
		}
		if !last {
			j, err = io.WriteString(w, "\n")
			n += j
		}
		return
	})
}

func (g *Generator) writeParagraph(w io.Writer) (int, error) {
	return writeRand(g.ParagraphLen, 1, func(_, _ bool) (n int, err error) {
		var j int
		j, err = g.writeSentence(w)
		n += j
		if err != nil {
			return
		}

		j, err = io.WriteString(w, "\n")
		n += j
		return
	})
}

func (g *Generator) writeSentence(w io.Writer) (int, error) {
	return writeRand(g.SentenceLen, 1, func(first, last bool) (n int, err error) {
		var j int
		if first {
			j, err = g.writeWordUpper(w)
			n += j
			if err != nil {
				return
			}
		}
		j, err = g.writeClause(w)
		n += j
		if err != nil {
			return
		}
		if last {
			j, err = io.WriteString(w, ".")
			n += j
		} else {
			m := len(g.Alphabet.Punct)
			i := g.SentenceLen.Rand.Intn(m)
			j, err = w.Write([]byte(g.Alphabet.Punct[i : i+1]))
			n += j
			j, err = io.WriteString(w, "\n")
			n += j
		}
		return
	})
}

func (g *Generator) writeClause(w io.Writer) (int, error) {
	return writeRand(g.ClauseLen, 1, func(_, last bool) (n int, err error) {
		var j int
		j, err = g.writeWord(w)
		n += j
		if err != nil {
			return
		}
		if !last {
			j, err = io.WriteString(w, " ")
			n += j
		}
		return
	})
}

func (g *Generator) writeWord(w io.Writer) (int, error)      { return io.WriteString(w, g.Word()) }
func (g *Generator) writeWordUpper(w io.Writer) (int, error) { return io.WriteString(w, g.WordUpper()) }

func (g *Generator) Document() string  { return usingBuffer(g.writeDocument) }
func (g *Generator) Paragraph() string { return usingBuffer(g.writeParagraph) }
func (g *Generator) Sentence() string  { return usingBuffer(g.writeSentence) }
func (g *Generator) Clause() string    { return usingBuffer(g.writeClause) }

func (g *Generator) Word() string {
	txt := word(g.WordLen)
	fill(g.WordLen, g.Alphabet.Lower, txt)
	return string(txt)
}

func (g *Generator) WordUpper() string {
	txt := word(g.WordLen)
	fill(g.WordLen, g.Alphabet.Upper, txt[0:1])
	fill(g.WordLen, g.Alphabet.Lower, txt[1:])
	return string(txt)
}

func writeRand(n *Normal, zero int, f func(first, last bool) (int, error)) (int, error) {
	sz := n.Sample()
	if sz == 0 {
		sz = zero
	}

	bw := 0
	for i := 0; i < sz; i++ {
		j, err := f(i == 0, i == sz-1)
		bw += j
		if err != nil {
			return bw, err
		}
	}
	return bw, nil
}

func usingBuffer(f func(w io.Writer) (int, error)) string {
	buf := &bytes.Buffer{}
	_, err := f(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func word(n *Normal) []byte {
	sz := n.Sample()
	if sz == 0 {
		sz = 1
	}
	txt := make([]byte, sz)
	return txt
}

func fill(n *Normal, from string, to []byte) {
	m := len(from)
	for i := range to {
		to[i] = from[n.Rand.Intn(m)]
	}
}
