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

const (
	staticDir = "/static/"
)

func main() {
	flag.Parse()
	http.Handle(staticDir, http.StripPrefix(staticDir, http.FileServer(http.Dir("../static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pusher, ok := w.(http.Pusher)
		if ok {
			if err := pusher.Push("/static/bootstrap.min.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/jquery-3.4.1.slim.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/popper.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/bootstrap.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/gopher.png", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		// fmt.Fprintf(w, http_http2.HTML)
		http_http2.Response(w)
	})
	log.Fatal(http.ListenAndServeTLS(*addr, "cert.pem", "key.pem", nil))
}
