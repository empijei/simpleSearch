package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/empijei/cli/lg"
)

var FastIndex = &FastSearcher{}

func main() {
	err := Load("data", Paragraphs, FastIndex)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	go View()
	openBrowser("http://localhost:42137")
	for m := range searchChannel {
		lg.Debug(m)
		t := time.Now()
		//p, err := Paragraphs.Search(m)
		p, err := FastIndex.Search(m)
		if err != nil {
			continue
		}
		lg.Infof("Lookup time: %d µs", time.Now().Sub(t).Nanoseconds()/1000)
		resultChannel <- p
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
