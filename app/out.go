package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/tpl"
	"net/http"
	"strings"

	//"net/url"

	//"strings"
)


func (o *OKNO) out(r *mux.Router) {
	////////////////
	// out
	////////////////
	a := r.Host("out.okno.rs").Subrouter()
	a.StrictSlash(true)

	a.HandleFunc("/", o.goodBye).Methods("GET")

	//a.Headers("Access-Control-Allow-Origin", "*")
}

func (o *OKNO) goodBye(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimSpace(r.URL.Query().Get("url"))
	img := strings.TrimSpace(r.URL.Query().Get("img"))
	title := strings.TrimSpace(r.URL.Query().Get("title"))
	out := map[string]string{
		"url":url,
		"img":img,
		"title":title,
	}
	tpl.TemplateHandler(cfg.Path+"/templates/okno").ExecuteTemplate(w, "out_gohtml", out)
}
