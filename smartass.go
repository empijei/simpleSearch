package main

/*
var Paragraphs ParagraphIndex

type ParagraphIndex struct {
	sync.Mutex
	index map[string]*Paragraph
}

func (p *ParagraphIndex) Search(title string) (*Paragraph, error) {
	p.Lock()
	defer p.Unlock()
	val, ok := p.index[title]
	if !ok {
		return nil, errors.New("Not found")
	}
	return val, nil
}

func (pi *ParagraphIndex) Add(c <-chan *Paragraph) {
	go func() {
		pi.Lock()
		for p := range c {
			pi.index[p.Title] = p
		}
		pi.Unlock()
	}()
}

//The following was commented out because it is part of the fast search algorithm
func (p *Paragraph) EngWords(c chan<- KeyValue) {
	sendWords(p.BodyEng, p.Title, c)
}

func (p *Paragraph) ItaWords(c chan<- KeyValue) {
	sendWords(p.BodyIta, p.Title, c)
}

func sendWords(content, title string, c chan<- KeyValue) {
	s := bufio.NewScanner(strings.NewReader(content))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		c <- KeyValue{s.Text(), title}
	}
	close(c)
}

type TitleSearcher struct {
	index wordIndex
}

func (t *TitleSearcher) Search(needle string) (*Paragraph, error) {
	panic("not implemented")
}

func (t *TitleSearcher) Add(*Paragraph) {
	panic("not implemented")
}
*/
