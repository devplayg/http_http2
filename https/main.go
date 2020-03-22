package main

import (
	"fmt"
	"github.com/devplayg/http_http2"
	"log"
	"net/http"
)

var addr = ":8001"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, http_http2.HTML)
	})
	log.Fatal(http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil))
}
