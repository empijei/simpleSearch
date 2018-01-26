package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func openToWrite(dir, title string) (io.WriteCloser, error) {
	filename := filepath.Join(dir, reg.ReplaceAllString(strings.Replace(strings.TrimSpace(strings.ToLower(title[2:])), " ", "_", -1), "")) + ".md"
	_ = os.Remove(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	return f, err
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Please provide a file")
		os.Exit(2)
	}
	inpath := os.Args[1]
	dir := filepath.Dir(inpath)
	f, err := os.Open(inpath)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	reg, _ = regexp.Compile("[^a-zA-Z0-9_]+")

	s := bufio.NewScanner(f)
	var w io.WriteCloser
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "# ") {
			if w != nil {
				_ = w.Close()
			}
			w, err = openToWrite(dir, s.Text())
			if err != nil {
				panic(err)
			}
		}
		if w != nil {
			fmt.Fprintln(w, s.Text())
		}
	}
	if w != nil {
		_ = w.Close()
	}
}
