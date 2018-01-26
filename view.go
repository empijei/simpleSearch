package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/empijei/cli/lg"
	"golang.org/x/net/websocket"
)

const entrytmpl = ` {{ range . }}
  <a href="#" class="list-group-item list-group-item-action flex-column align-items-start active" onclick="listclick(event);">
    <div class="d-flex w-100 justify-content-between">
      <h5 class="mb-1">{{.Title}}</h5>
		<small>{{.Classification}}</small>
    </div>
    <p class="mb-1">{{.BodyEng}}</p>
	 <br>
    <p class="mb-1">{{.BodyIta}}</p>
	 <small>{{.Activity}}</small>
  </a>
{{end}}`

var pentrytmpl *template.Template

func View() {
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	onConnected := func(ws *websocket.Conn) {
		lg.Info("A client has connected")
		handleClient(ws)
	}
	pentrytmpl = template.Must(template.New("entry").Parse(entrytmpl))
	http.Handle("/ws", websocket.Handler(onConnected))
	lg.Info("Starting web interface")
	_ = http.ListenAndServe("127.0.0.1:42137", nil)
}

var searchChannel = make(chan string)
var resultChannel = make(chan []*Paragraph)

func parToView(pars []*Paragraph) map[string]string {
	buf := bytes.NewBuffer(nil)
	err := pentrytmpl.Execute(buf, pars)
	if err != nil {
		lg.Error(err)
	}
	return map[string]string{"Html": string(buf.Bytes())}
}

func handleClient(ws io.ReadWriteCloser) {
	defer func() { _ = ws.Close(); lg.Info("A client has disconnected") }()
	go func() {
		defer func() { _ = ws.Close() }()
		e := json.NewEncoder(ws)
		for p := range resultChannel {
			err := e.Encode(parToView(p))
			_ = p
			if err != nil {
				lg.Error(err)
				return
			}
		}
	}()
	d := json.NewDecoder(ws)
	msg := make(map[string]string)
	var err error
	for {
		err = d.Decode(&msg)
		if err != nil {
			lg.Error(err)
			return
		}
		lg.Debug(msg)
		searchChannel <- msg["search"]
	}
}
