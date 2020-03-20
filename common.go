package http_http2

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	StaticDir = "/static/"
)

func Response(w http.ResponseWriter) {
	vars := map[string]interface{}{
		"ts":  time.Now().UnixNano(),
		"arr": make([]int, 36),
	}

	html, _ := ioutil.ReadFile("../index.html")
	t := template.Must(template.New("html").Parse(string(html)))
	t.Execute(w, vars)
}
