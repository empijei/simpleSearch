package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/empijei/cli/lg"
)

//TODO make this a vararg of wheres and add support for streamers
func Load(path string, where ...Searcher) error {
	lg.Info("Recomputing indexes...")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, w := range where {
		if w, ok := w.(StreamSearcher); ok {
			w.Open()
		}
	}
	t := time.Now()
	i := 0
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			f, err := os.Open(filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
			i++
			p, err := LoadParagraph(f)
			if err != nil {
				return err
			}
			for _, w := range where {
				w.Add(p)
			}
		}
	}
	for _, w := range where {
		if w, ok := w.(StreamSearcher); ok {
			w.Close()
		}
	}
	lg.Infof("Indexes updated, %d files added in %d ms", i, time.Now().Sub(t).Nanoseconds()/1000000)
	return nil
}

func Save(path string, what Searcher, filename string) error {
	filename = filepath.Join(path, filename)
	filename += ".index"
	_ = os.Remove(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	e := json.NewEncoder(f)
	err = e.Encode(what)
	if err != nil {
		return err
	}
	return nil
}
