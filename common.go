package http_http2

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

func Response(w http.ResponseWriter) {
	vars := map[string]interface{}{
		"ts":  time.Now().UnixNano(),
		"arr": make([]int, 60),
	}

	html, _ := ioutil.ReadFile("../index.html")
	t := template.Must(template.New("html-tmpl").Parse(string(html)))
	t.Execute(w, vars)
}
