package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Load(path string, where Searcher) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			f, err := os.Open(filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
			p, err := LoadParagraph(f)
			if err != nil {
				return err
			}
			where.Add(p)
		}
	}
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
