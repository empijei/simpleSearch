package main

type ParSet map[*Paragraph]struct{}

func (ps ParSet) Add(p *Paragraph) {
	if ps == nil {
		ps = make(ParSet)
	}
	ps[p] = struct{}{}
}

func (ps ParSet) GetSlice() (slice []*Paragraph) {
	for p := range ps {
		slice = append(slice, p)
	}
	return
}

func (ps1 ParSet) Intersection(ps2 ParSet) (i ParSet) {
	if len(ps2) < len(ps1) {
		ps2, ps1 = ps1, ps2
	}
	for p := range ps1 {
		if _, ok := ps2[p]; ok {
			i.Add(p)
		}
	}
	return
}
