package main

import "log"

func main() {
	err := Load("data", Paragraphs)
	if err != nil {
		log.Println(err)
	}
	err = Save("data", Paragraphs, "SlowSearch")
	if err != nil {
		log.Println(err)
	}
}
