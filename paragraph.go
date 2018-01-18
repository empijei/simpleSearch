package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

var Paragraphs = &SlowSearcher{}

type Paragraph struct {
	Title, BodyEng, BodyIta string
}

func LoadParagraph(r io.Reader) (p *Paragraph, err error) {
	s := bufio.NewScanner(r)
	p = &Paragraph{}
	var line string
	itabuf := bytes.NewBuffer(nil)
	engbuf := bytes.NewBuffer(nil)
	var isIta bool
	for s.Scan() {
		line = s.Text()
		if strings.HasPrefix(line, "# ") {
			p.Title = line[2:]
			continue
		}
		if strings.HasPrefix(line, "## E") {
			isIta = false
			continue
		}
		if strings.HasPrefix(line, "## I") {
			isIta = true
			continue
		}
		toAdd := line
		if len(line) == 0 {
			toAdd += "\n"
		}
		if isIta {
			itabuf.WriteString(toAdd)
		} else {
			engbuf.WriteString(toAdd)
		}
	}
	p.BodyEng = string(engbuf.Bytes())
	p.BodyIta = string(itabuf.Bytes())
	return p, nil
}
