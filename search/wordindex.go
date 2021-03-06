package main

import (
	"bufio"
	"strings"
	"sync"

	"github.com/empijei/cli/lg"
)

//All of the following constants are just arbitrary
const MaxResultSize = 10
const MinSearchLength = 3
const chanSize = 10

type wordIndex struct {
	index map[string]*ParSet
}

func (w *wordIndex) lookup(needle string) *ParSet {
	return w.index[strings.ToLower(needle)]
}

//This is thread-unsafe, it expects the caller to perform checks
func (wi *wordIndex) addToIndex(c <-chan *Paragraph) {
	wi.index = make(map[string]*ParSet)
	for p := range c {
		for w := range getWordsStream(strings.ToLower(p.GetAllText())) {
			//lg.Debugf("Inserting <%v> into index", w)
			for _, sw := range getSubWords(w) {
				if wi.index[sw] == nil {
					wi.index[sw] = &ParSet{}
				}
				wi.index[sw].Add(p)
			}

		}
	}
}

func getSubWords(word string) (subwords []string) {
	if len(word) < MinSearchLength {
		return
	}
	for size := MinSearchLength; size <= len(word); size++ {
		for lower := 0; lower <= len(word)-size; lower++ {
			subwords = append(subwords, word[lower:lower+size])
		}
	}
	return
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
		if len(s.Text()) >= MinSearchLength {
			splitted = append(splitted, s.Text())
		}
	}
	return
}

type FastSearcher struct {
	sync.Mutex

	Index *wordIndex

	done  chan struct{}
	learn chan *Paragraph
}

func (fs *FastSearcher) Search(needle string) ([]*Paragraph, error) {
	//lg.Debug("Searching for: ", needle)
	words := getWords(needle)
	if len(words) == 0 {
		return nil, nil
	}
	set := fs.Index.lookup(words[0])
	for i := 1; i < len(words); i++ {
		if set == nil {
			return nil, nil
		}
		set = set.Intersection(fs.Index.lookup(words[i]))
	}
	if set == nil {
		return nil, nil
	}
	return set.GetCroppedSlice(MaxResultSize), nil
}

func (fs *FastSearcher) Add(p *Paragraph) {
	//lg.Debug("Adding paragraph ", p.Title)
	fs.Lock()
	defer func() {
		fs.Unlock()
		if err := recover(); err != nil {
			lg.Error(err)
		}
	}()
	if fs.learn == nil {
		lg.Error("Adding not in learn mode")
		return
	}
	fs.learn <- p
}

func (fs *FastSearcher) Open() {
	//lg.Debug("Started learning")
	fs.Lock()
	defer fs.Unlock()
	if fs.learn == nil {
		fs.learn = make(chan *Paragraph, chanSize)
		fs.done = make(chan struct{})
	}
	go func() {
		fs.Index = &wordIndex{}
		fs.Index.addToIndex(fs.learn)
		fs.Lock()
		fs.learn = nil
		fs.Unlock()
		//lg.Debug("Learning phase closed")
		fs.done <- struct{}{}
	}()
}

func (fs *FastSearcher) Close() {
	//lg.Debug("Closing learning phase")
	close(fs.learn)
	<-fs.done
	//lg.Debug("Learnt: ", *fs.Index)
}
