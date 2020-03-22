package main

import (
	"fmt"
	"github.com/devplayg/http_http2"
	"log"
	"net/http"
)

var addr = ":8002"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pusher, ok := w.(http.Pusher)
		if ok {
			push(pusher, http_http2.Assets, nil)

		}
		fmt.Fprintf(w, http_http2.HTML)
	})

	// tlsConfig := &tls.Config{
	// 	MinVersion:               tls.VersionTLS12,
	// 	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 	PreferServerCipherSuites: true,
	// 	CipherSuites: []uint16{
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	// 	},
	// }

	// srv := &http.Server{
	// 	Addr:        addr,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	//
	// srv := &http.Server{
	// 	Addr: addr,
	// 	// TLSConfig:tlsConfig,
	// 	// TLSConfig:    tlsConfig,
	// 	ReadTimeout:       1 * time.Second,
	// 	WriteTimeout:      1 * time.Second,
	// 	IdleTimeout:       30 * time.Second,
	// 	ReadHeaderTimeout: 2 * time.Second,
	// }
	// log.Fatal(srv.ListenAndServeTLS("cert.crt", "key.key"))
	// log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))

	log.Fatal(http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil))
}

func push(pusher http.Pusher, paths []string, options *http.PushOptions) {

	for _, p := range paths {
		if err := pusher.Push(p, options); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
}
