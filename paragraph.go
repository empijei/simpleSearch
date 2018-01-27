package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/empijei/cli/lg"
)

type Paragraph struct {
	Title, BodyEng, BodyIta, Activity, Classification, Score string
}

func LoadParagraph(r io.Reader) (p *Paragraph, err error) {
	s := bufio.NewScanner(r)
	p = &Paragraph{}
	var line string
	itabuf := bytes.NewBuffer(nil)
	engbuf := bytes.NewBuffer(nil)
	var isIta bool
	isMeta := true
	for s.Scan() {
		line = s.Text()
		if strings.HasPrefix(line, "# ") {
			p.Title = line[2:]
			continue
		}
		if isMeta {
			switch {
			case strings.HasPrefix(line, "Score:"):
				p.Score = line[6:]
			case strings.HasPrefix(line, "Activ:"):
				p.Activity = line[6:]
			case strings.HasPrefix(line, "OWASP:"):
				p.Classification = line[6:]
			default:
				if len(line) > 0 && line[0] != byte('#') {
					lg.Errorf("Unknown meta: <%s> in Paragraph <%s>", line, p.Title)
				}
			}
			if !strings.HasPrefix(line, "## ") {
				continue
			}
		}
		if strings.HasPrefix(line, "## E") {
			isIta = false
			isMeta = false
			continue
		}
		if strings.HasPrefix(line, "## I") {
			isIta = true
			isMeta = false
			continue
		}
		toAdd := line
		if isIta {
			itabuf.WriteString(toAdd)
			itabuf.WriteRune('\n')
		} else {
			engbuf.WriteString(toAdd)
			engbuf.WriteRune('\n')
		}
	}
	p.BodyEng = string(engbuf.Bytes())
	p.BodyIta = string(itabuf.Bytes())
	return p, nil
}

func (p *Paragraph) GetAllText() string {
	return p.Title + " " + p.BodyIta + " " + p.BodyEng + " " + p.Activity + " " + p.Classification + " " + p.Score
}
