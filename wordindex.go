package main

import (
	"bufio"
	"strings"
	"sync"

	"github.com/empijei/wapty/cli/lg"
)

type wordIndex struct {
	index map[string]ParSet
}

func (w *wordIndex) lookup(needle string) ParSet {
	return w.index[needle]
}

//This is thread-unsafe, it expects the caller to perform checks
func (wi *wordIndex) addToIndex(c <-chan *Paragraph) {
	for p := range c {
		for w := range getWordsStream(p.GetAllText()) {
			wi.index[w].Add(p)
		}
	}
}

func getWordsStream(words string) <-chan string {
	c := make(chan string, 15)
	go func() {
		s := bufio.NewScanner(strings.NewReader(words))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()
	return c
}

func getWords(words string) (splitted []string) {
	s := bufio.NewScanner(strings.NewReader(words))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		splitted = append(splitted, s.Text())
	}
	return
}

type FastSearcher struct {
	sync.Mutex

	index wordIndex
	learn chan *Paragraph
}

func (fs *FastSearcher) Search(needle string) ([]*Paragraph, error) {
	words := getWords(needle)
	if len(words) == 0 {
		return nil, nil
	}
	set := fs.index.lookup(words[0])
	for i := 1; i < len(words); i++ {
		set = set.Intersection(fs.index.lookup(words[i]))
	}
	return set.GetSlice(), nil
}

func (fs *FastSearcher) Add(p *Paragraph) {
	fs.Lock()
	defer func() {
		fs.Unlock()
		if err := recover(); err != nil {
			lg.Error(err)
		}
	}()
	if fs.learn == nil {
		return
	}
	fs.learn <- p
}

func (fs *FastSearcher) Open() {
	fs.Lock()
	defer fs.Unlock()
	if fs.learn == nil {
		fs.learn = make(chan *Paragraph)
	}
	go func() {
		fs.index.addToIndex(fs.learn)
		fs.Lock()
		fs.learn = nil
		fs.Unlock()
	}()
}

func (fs *FastSearcher) Close() {
	close(fs.learn)
}
