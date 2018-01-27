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
*/

type TitleSearcher struct {
	pars map[string][]*Paragraph
}

func (t *TitleSearcher) Search(title string) ([]*Paragraph, error) {
	p, ok := t.pars[title]
	if !ok {
		return nil, NotFoundError
	}
	return p, nil
}

func (t *TitleSearcher) Add(p *Paragraph) {
	t.pars[p.Title] = append(t.pars[p.Title], p)
}
func (t *TitleSearcher) Open() {
	t.pars = make(map[string][]*Paragraph)
}
func (t *TitleSearcher) Close() {}
