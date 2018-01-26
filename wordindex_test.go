package main

import (
	"reflect"
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
		{"tIt11", []string{"Tit1"}},
		{"tit21", []string{"Tit2"}},
		{"act", []string{"Tit2", "Tit1"}},
		{"", nil},
		{"tit", nil},
	}

	for i, tt := range lookuptest {
		res := wi.lookup(tt.in)
		if len(tt.out) != len(res.GetSlice()) {
			t.Error("TestAddToIndex[%d]", i)
			continue
		}
		//TODO
	}
}
