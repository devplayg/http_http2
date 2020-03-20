package main

import (
	"flag"
	"github.com/devplayg/http_http2"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8000", "Listen address")
)

func main() {
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http_http2.Response(w)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
