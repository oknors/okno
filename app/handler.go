package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/mod"
	"github.com/oknors/okno/app/tpl"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"os"
)

type PageData struct {
	Host Host
	Post mod.Post

	Posts []mod.Post
	Hosts []Host
}

func (o *OKNO) Handler() http.Handler {
	r := mux.NewRouter()

	for _, h := range o.Hosts {
		sh := h.sub(r)
		o.index(sh, h)
		o.post(sh, h)
		o.staticHost(sh, h.Slug)

	}
	o.oknoAdmin(r)
	o.static(r)

	o.explorer(r)

	//o.api(r)
	o.jorm(r)
	//o.mattermost(r)
	o.wing(r)
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

// HomeHandler handles a request for (?)
func (o *OKNO) index(rt *mux.Router, host Host) {
	rt.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &PageData{
			Host:  host,
			Posts: o.posts("/sites/" + host.Slug + "/jdb/" + "posts"),
		}
		templatePath := o.Configuration.Path + "/sites/" + host.Slug + "/tpl/gohtml/index.gohtml"
		_, err := os.Stat(templatePath)
		if err != nil {
			tpl.TemplateHandler(o.Configuration.Path+"/sites").ExecuteTemplate(w, "err_gohtml", err)
		} else {
			tpl.TemplateHandler(o.Configuration.Path+"/sites/"+host.Slug).ExecuteTemplate(w, "index_gohtml", data)
		}
	})
}

// HomeHandler handles a request for (?)
func (o *OKNO) post(rt *mux.Router, host Host) {
	rt.HandleFunc("/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {
		//col := mux.Vars(r)["col"]
		id := mux.Vars(r)["slug"]
		post := o.Database.ReadPost("/sites/"+host.Slug+"/jdb", "posts", id)
		data := &PageData{
			Host: host,
			Post: post,
		}
		tpl.TemplateHandler(o.Configuration.Path+"/sites/"+host.Slug).ExecuteTemplate(w, "post_gohtml", data)
	})
}

func (o *OKNO) static(r *mux.Router) {
	s := r.Host("s.okno.rs").Subrouter()
	s.StrictSlash(true).PathPrefix("/").Handler(http.FileServer(http.Dir(o.Configuration.Path + "/static")))
}

func (o *OKNO) staticHost(r *mux.Router, host string) {
	p := o.Configuration.Path + "/sites/" + host + "/static"
	//r.StrictSlash(true).
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir( p)))
	r.StrictSlash(true).PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(p))))
}
