package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetWordsAndStream(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			"",
			nil,
		},
		{
			"prova",
			[]string{"prova"},
		},
		{
			"prova ",
			[]string{"prova"},
		},
		{
			"prova uno",
			[]string{"prova", "uno"},
		},
		{
			`prova
due`,
			[]string{"prova", "due"},
		},
		{
			"prova         due",
			[]string{"prova", "due"},
		},
	}
	for i, tt := range tests {
		out := getWords(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("%d: getWords(\"%s\") expected: <%#v> got: <%#v>", i, tt.in, tt.out, out)
		}
	}

	for i, tt := range tests {
		outc := getWordsStream(tt.in)
		var out []string
		for w := range outc {
			out = append(out, w)
		}
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("%d: getWordsStream(\"%s\") expected: <%#v> got: <%#v>", i, tt.in, tt.out, out)
		}
	}

}

func TestAddToIndex(t *testing.T) {
	wi := &wordIndex{}
	c := make(chan *Paragraph, 3)
	c <- &Paragraph{
		"Tit11 tIt12",
		"Eng11 eNg12",
		"Ita11 iTa12",
		"Act aCt",
		"Cla11 cLa12",
		"Sco11 sCo12",
	}
	c <- &Paragraph{
		"Tit21 tIt22",
		"Eng21 eNg22",
		"Ita21 iTa22",
		"Act aCt",
		"Cla21 cLa22",
		"Sco21 sCo22",
	}
	close(c)

	wi.addToIndex(c)

	lookuptest := []struct {
		in  string
		out []string
	}{
		{"tIt11", []string{"Tit11 tIt12"}},
		{"tit21", []string{"Tit21 tIt22"}},
		{"act", []string{"Tit21 tIt22", "Tit11 tIt12"}},
		{"", nil},
		{"tit", []string{"Tit21 tIt22", "Tit11 tIt12"}},
	}

	getParTitles := func(pars []*Paragraph) []string {
		var titles []string
		for _, p := range pars {
			titles = append(titles, p.Title)
		}
		return titles
	}

	for i, tt := range lookuptest {
		res := wi.lookup(tt.in)
		if tt.out == nil && res == nil {
			continue
		}
		tits := getParTitles(res.GetSlice())
		if areDiffStringSlices(tits, tt.out) {
			t.Errorf("TestAddToIndex[%d], expected <%#v> got <%#v>", i, tt.out, tits)
			continue
		}
	}
}

func TestGetSubWords(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			"word",
			[]string{"wor", "ord", "word"},
		},
		{
			"in",
			nil,
		},
		{
			"longword",
			[]string{"lon", "ong", "ngw", "gwo", "wor", "ord", "long", "ongw", "ngwo", "gwor", "word", "longw", "ongwo", "ngwor", "gword", "longwo", "ongwor", "ngword", "longwor", "ongword", "longword"},
		},
		{
			"wor",
			[]string{"wor"},
		},
	}

	for i, tt := range tests {
		out := getSubWords(tt.in)
		if areDiffStringSlices(tt.out, out) {
			t.Errorf("%d: getSubWords(\"%s\") expected: <%#v> got: <%#v>", i, tt.in, tt.out, out)
		}
	}
}

func areDiffStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return true
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, w := range a {
		if w != b[i] {
			return true
		}
	}
	return false
}
