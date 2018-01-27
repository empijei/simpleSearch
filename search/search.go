package main

import (
	"errors"
)

var NotFoundError = errors.New("Not Found")

type Searcher interface {
	Search(needle string) (found []*Paragraph, err error)
	Add(p *Paragraph)
}

type StreamSearcher interface {
	Searcher
	Open()
	Close()
}
