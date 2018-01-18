package main

import (
	"errors"
	"strings"
	"sync"
)

var NotFoundError = errors.New("Not Found")

type Searcher interface {
	Search(needle string) (found []*Paragraph, err error)
	Add(p *Paragraph)
}

type SlowSearcher struct {
	//filter FilterFlag
	sync.Mutex
	Pars []*Paragraph
}

func (s *SlowSearcher) Search(needle string) (found []*Paragraph, err error) {
	for _, p := range s.Pars {
		if strings.Contains(p.Title, needle) {
			found = append(found, p)
		}
		if strings.Contains(p.BodyEng, needle) {
			found = append(found, p)
		}
		if strings.Contains(p.BodyIta, needle) {
			found = append(found, p)
		}
	}
	if found == nil {
		err = NotFoundError
	}
	return
}

func (s *SlowSearcher) Add(p *Paragraph) {
	s.Lock()
	s.Pars = append(s.Pars, p)
	s.Unlock()
}
