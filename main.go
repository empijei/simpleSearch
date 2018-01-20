package main

import (
	"log"
	"time"

	"github.com/empijei/wapty/cli/lg"
)

var FastIndex = &FastSearcher{}

func main() {
	err := Load("data", Paragraphs, FastIndex)
	if err != nil {
		log.Println(err)
	}
	err = Save("data", Paragraphs, "SlowSearch")
	err = Save("data", FastIndex, "FastSearch")
	if err != nil {
		log.Println(err)
	}
	go func() {
		for m := range searchChannel {
			lg.Debug(m)
			//p, err := Paragraphs.Search(m)
			t := time.Now()
			p, err := FastIndex.Search(m)
			if err != nil {
				continue
			}
			lg.Infof("Lookup time: %d ms", time.Now().Sub(t).Nanoseconds()/1000000)
			resultChannel <- p
		}
	}()
	View()
}
