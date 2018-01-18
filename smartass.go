package main

/*
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
*/

/*
//The following was commented out because it is part of the fast search algorithm
type wordIndex struct {
	sync.Mutex

	index map[string][]string
}

func (w *wordIndex) lookup(needle string) []string {
	return w.index[needle]
}

type KeyValue struct {
	//Key will be a word in a paragraph, value will be the title of the paragraph
	Key, Value string
}

func (wi *wordIndex) addToIndex(c <-chan KeyValue) {
	go func() {
		wi.Lock()
		for wt := range c {
			wi.index[wt.Key] = append(wi.index[wt.Key], wt.Value)
		}
		wi.Unlock()
	}()
}

type BodySearcher struct {
	index wordIndex
}

func (b *BodySearcher) Search(needle string) (*Paragraph, error) {
	panic("not implemented")
}

func (b *BodySearcher) Add(*Paragraph) {
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
