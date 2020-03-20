package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	addr = flag.String("addr", ":8000", "Listen address")
)

//func Response(w http.ResponseWriter) {
//	vars := map[string]interface{} {
//		"ts":  time.Now().UnixNano(),
//		"arr": make([]int, 60),
//	}
//
//	html, _ := ioutil.ReadFile("../index.html")
//	t := template.Must(template.New("html-tmpl").Parse(string(html)))
//	t.Execute(w, vars)
//}

func main() {
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Response(w)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
