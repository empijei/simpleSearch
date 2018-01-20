package main

type ParSet struct {
	s map[*Paragraph]struct{}
}

func (ps *ParSet) Add(p *Paragraph) {
	if ps.s == nil {
		ps.s = make(map[*Paragraph]struct{})
	}
	ps.s[p] = struct{}{}
}

func (ps ParSet) GetSlice() (slice []*Paragraph) {
	for p := range ps.s {
		slice = append(slice, p)
	}
	return
}
func (ps ParSet) GetCroppedSlice(count int) (slice []*Paragraph) {
	i := 0
	for p := range ps.s {
		slice = append(slice, p)
		i++
		if i > count {
			return
		}
	}
	return
}

func (ps1 ParSet) Intersection(ps2 ParSet) (i *ParSet) {
	if len(ps2.s) < len(ps1.s) {
		ps2, ps1 = ps1, ps2
	}
	i = &ParSet{}
	for p := range ps1.s {
		if _, ok := ps2.s[p]; ok {
			i.Add(p)
		}
	}
	return
}
