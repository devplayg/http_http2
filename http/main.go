package main

import (
	"flag"
	"fmt"
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
		fmt.Fprintf(w, http_http2.HTML)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
