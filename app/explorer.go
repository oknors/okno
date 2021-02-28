package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/tpl"

	"net/http"
)

func (o *OKNO) explorer(r *mux.Router) {
	////////////////
	// explorer
	////////////////
	s := r.Host("explorer.parallelcoin.info").Subrouter()
	s.StrictSlash(true)
	s.HandleFunc("/", o.explorerIndex)
	s.HandleFunc("/{section}", o.explorerSection)
	s.HandleFunc("/{type}/{slug}", o.explorerItem)
}

// HomeHandler handles a request for (?)
func (o *OKNO) explorerIndex(w http.ResponseWriter, r *http.Request) {
	tpl.TemplateHandler(o.Configuration.Path+"/sites/explorer_parallelcoin_info").ExecuteTemplate(w, "index_gohtml", nil)
}

// HomeHandler handles a request for (?)
func (o *OKNO) explorerSection(w http.ResponseWriter, r *http.Request) {
	tpl.TemplateHandler(o.Configuration.Path+"/sites/explorer_parallelcoin_info").ExecuteTemplate(w, "section_gohtml", mux.Vars(r)["section"])
}

// HomeHandler handles a request for (?)
func (o *OKNO) explorerItem(w http.ResponseWriter, r *http.Request) {
	data := struct {
		T    string
		Slug string
	}{mux.Vars(r)["type"], mux.Vars(r)["slug"]}

	tpl.TemplateHandler(o.Configuration.Path+"/sites/explorer_parallelcoin_info").ExecuteTemplate(w, "item_gohtml", data)
}
