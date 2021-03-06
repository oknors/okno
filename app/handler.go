package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/mod"
	"github.com/oknors/okno/app/tpl"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"os"
	"text/template"
)

type PageData struct {
	Title   string
	Host    Host
	Site    string
	Post    mod.Post
	Slug    string
	Section string
	Posts   []mod.Post
	Hosts   []Host
}

func (o *OKNO) Handler() http.Handler {
	r := mux.NewRouter()

	for _, h := range o.Hosts {
		dh := h.domain(r)
		o.index(dh, h)
		//o.section(dh, h)
		o.post(dh, h)
		//o.staticHost(dh, h.Slug)
		if h.Slug != "okno_rs" {
			sh := h.sub(r)
			o.index(sh, h)
			//o.section(sh, h)
			o.post(sh, h)
			//o.staticHost(sh, h.Slug)
		}
	}
	o.oknoAdmin(r)
	o.static(r)

	//o.explorer(r)

	o.api(r)
	o.img(r)
	o.jorm(r)
	//o.mattermost(r)
	o.wing(r)
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

// HomeHandler handles a request for (?)
func (o *OKNO) index(rt *mux.Router, host Host) {
	rt.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		site := mux.Vars(r)["site"]
		title := host.Name
		hostSlug := host.Slug
		if site != "" {
			title = site + " " + host.Name
			hostSlug = site + "_" + host.Slug
		}
		data := &PageData{
			Title:   title,
			Host:    host,
			Site:    hostSlug,
			Section: "index",
			//Posts: o.posts("/sites/" + hostSlug + "/jdb/" + "post"),
		}
		o.template(w, host, data)
	})
}

// HomeHandler handles a request for (?)
func (o *OKNO) section(rt *mux.Router, host Host) {
	rt.HandleFunc("/{section}", func(w http.ResponseWriter, r *http.Request) {
		site := mux.Vars(r)["site"]
		section := mux.Vars(r)["section"]
		title := host.Name
		hostSlug := host.Slug
		if site != "" {
			title = site + " " + host.Name
			hostSlug = site + "_" + host.Slug
		}
		data := &PageData{
			Title:   title,
			Host:    host,
			Site:    hostSlug,
			Section: section,
			//Posts: o.posts("/sites/" + hostSlug + "/jdb/" + section),
		}
		o.template(w, host, data)
	})
}

// HomeHandler handles a request for (?)
func (o *OKNO) post(rt *mux.Router, host Host) {
	rt.HandleFunc("/{section}/{slug}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["slug"]
		site := mux.Vars(r)["site"]
		section := mux.Vars(r)["section"]
		hostSlug := host.Slug
		if site != "" {
			hostSlug = site + "_" + host.Slug
		}
		//post := o.Database.ReadPost("/sites/"+ hostSlug +"/jdb", section, id)
		data := &PageData{
			Host: host,
			Site: hostSlug,
			//Post: post,
			Slug:    id,
			Section: section,
		}
		o.template(w, host, data)
	})
}

func (o *OKNO) template(w http.ResponseWriter, host Host, data interface{}) {
	funcMap := template.FuncMap{
		"truncate": truncate,
	}
	templatePath := o.Configuration.Path + "/sites/" + host.Slug + "/tpl/gohtml/index.gohtml"
	if host.Template != "" {
		templatePath = o.Configuration.Path + "/templates/" + host.Template + "/tpl/gohtml/index.gohtml"
	}
	_, err := os.Stat(templatePath)
	if err != nil {
		tpl.TemplateHandler(o.Configuration.Path+"/sites").Funcs(funcMap).ExecuteTemplate(w, "err_gohtml", err)
	} else {
		if host.Template != "" {
			tpl.TemplateHandler(o.Configuration.Path+"/templates/"+host.Template).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		} else {
			tpl.TemplateHandler(o.Configuration.Path+"/sites/"+host.Slug).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		}
	}
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

func truncate(s string, l int) string {
	if l != 0 {
		var numRunes = 0
		for index, _ := range s {
			numRunes++
			if numRunes > l {
				return s[:index]
			}
		}
	}
	return s
}
