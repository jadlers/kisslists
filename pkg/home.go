package pkg

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var html, _ = ioutil.ReadFile("static/frontend.html")
var frontend = string(html)
var homeTempl = template.Must(template.New("").Parse(frontend))

func ServeHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var v = struct {
			Host string
		}{
			r.Host,
		}
		err := homeTempl.Execute(w, &v)
		if err != nil {
			log.Println(err)
		}
	}
}
