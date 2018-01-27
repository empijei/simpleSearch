package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/empijei/cli/lg"
)

var FastIndex = &FastSearcher{}
var SlowSearch = &SlowSearcher{}
var TitleIndex = &TitleSearcher{}

var selected []*Paragraph

func main() {
	err := Load("data", SlowSearch, FastIndex, TitleIndex)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	go View()
	openBrowser("http://localhost:42137")
	go searchJob()
	addSelected()
}

func addSelected() {
	for m := range selectChan {
		res, err := TitleIndex.Search(m)
		if err != nil {
			lg.Error(err, m)
			continue
		}
		selected = append(selected, res...)
		lg.Debugf("Added paragraph \"%s\", total selected: %d", res[0].Title, len(selected))
		resultChan <- map[string]string{"Result": strconv.Itoa(len(selected))}

	}
}

func searchJob() {
	for m := range searchChan {
		lg.Debug(m)
		t := time.Now()
		//p, err := Paragraphs.Search(m)
		p, err := FastIndex.Search(m)
		if err != nil {
			continue
		}
		lg.Infof("Lookup time: %d µs", time.Now().Sub(t).Nanoseconds()/1000)
		resultChan <- p
	}
}

func openBrowser(url string) {
	fmt.Println("Visiting: " + url)
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("start", url)
	}
	var err error
	if cmd != nil {
		err = cmd.Start()
		if err != nil {
			lg.Error(err)
		}
	} else {
		lg.Error("Unknown os: ", runtime.GOOS)
	}
	if err != nil || cmd == nil {
		fmt.Println("Cannot open browser, please visit ", url)
	}
}
