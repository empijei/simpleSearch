package main

import (
	"log"

	"github.com/empijei/wapty/cli/lg"
)

func main() {
	err := Load("data", Paragraphs)
	if err != nil {
		log.Println(err)
	}
	err = Save("data", Paragraphs, "SlowSearch")
	if err != nil {
		log.Println(err)
	}
	go func() {
		for m := range searchChannel {
			lg.Debug(m)
		}
	}()
	View()
}
