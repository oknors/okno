package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/mod"
	"github.com/oknors/okno/app/tpl"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"os"
	"text/template"
)

type PageData struct {
	Title    string
	Host     Host
	Site     string
	HostSlug string
	Post     mod.Post
	Slug     string
	Section  string
	Template string
	Posts    []mod.Post
	Hosts    []Host
}

func (o *OKNO) Handler() http.Handler {
	r := mux.NewRouter()

	for _, h := range o.Hosts {
		dh := h.domain(r)
		o.index(dh, h)
		//o.section(dh, h)
		o.post(dh, h)
		//o.staticHost(dh, h.Slug)
		if h.Slug != "okno_rs" && h.Slug != "marcetin_com" {
			sh := h.sub(r)
			o.index(sh, h)
			//o.section(sh, h)
			o.post(sh, h)
			//o.staticHost(sh, h.Slug)
			//if h.Template != "parallelcoin" {
			//	o.chat(sh)
			//}
		}
	}
	o.oknoAdmin(r)
	o.static(r)
	o.templates(r)

	o.api(r)
	o.img(r)
	o.out(r)
	o.jorm(r)
	o.our(r)
	o.wing(r)
	return handlers.CORS()(handlers.CompressHandler(InterceptHandler(r, DefaultErrorHandler)))
}

// HomeHandler handles a request for (?)
func (o *OKNO) index(rt *mux.Router, host Host) {
	rt.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		site := mux.Vars(r)["site"]
		title := host.Name
		hostSlug := host.Slug
		section := "index"
		template := site
		if site != "" {
			title = site + " " + host.Name
			hostSlug = site + "_" + host.Slug
			section = site
		} else {
			site = hostSlug
		}
		data := &PageData{
			Title:    title,
			Host:     host,
			HostSlug: hostSlug,
			Site:     site,
			Section:  section,
			Template: template,
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
		template := site + "/" + section
		if site != "" {
			title = site + " " + host.Name
			hostSlug = site + "_" + host.Slug
		}
		data := &PageData{
			Title:   title,
			Host:    host,
			Site:    hostSlug,
			Section: section,
			Template: template,
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
		template := site
		hostSlug := host.Slug
		if site != "" {
			hostSlug = site + "_" + host.Slug
			template = site + "/" + section
		}
		//post := o.Database.ReadPost("/sites/"+ hostSlug +"/jdb", section, id)
		data := &PageData{
			Host:     host,
			HostSlug: hostSlug,
			Site:     site,
			//Post: post,
			Template: template,
			Slug:    id,
			Section: section,
		}
		o.template(w, host, data)
	})
}

func (o *OKNO) template(w http.ResponseWriter, host Host, data interface{}) {
	funcMap := template.FuncMap{
		"truncate": utl.Truncate,
		"sha384": utl.SHA384,
	}
	templatePath := cfg.Path + "/sites/" + host.Slug + "/tpl/gohtml/index.gohtml"
	if host.Template != "" {
		templatePath = cfg.Path + "/templates/" + host.Template + "/tpl/gohtml/index.gohtml"
	}
	_, err := os.Stat(templatePath)
	if err != nil {
		tpl.TemplateHandler(cfg.Path+"/sites").Funcs(funcMap).ExecuteTemplate(w, "err_gohtml", err)
	} else {
		if host.Template != "" {
			tpl.TemplateHandler(cfg.Path+"/templates/"+host.Template).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		} else {
			tpl.TemplateHandler(cfg.Path+"/sites/"+host.Slug).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		}
	}
}

func (o *OKNO) static(r *mux.Router) {
	s := r.Host("s.okno.rs").Subrouter()
	s.StrictSlash(true).PathPrefix("/").Handler(http.FileServer(http.Dir(cfg.Path + "/static")))
	s.Headers("Access-Control-Allow-Headers", "*")
	s.Headers("Access-Control-Allow-Origin", "*")


}
func (o *OKNO) templates(r *mux.Router) {
	s := r.Host("t.okno.rs").Subrouter()
	s.StrictSlash(true).PathPrefix("/").Handler(http.FileServer(http.Dir(cfg.Path + "/templates")))
	s.Headers("Access-Control-Allow-Headers", "*")
	s.Headers("Access-Control-Allow-Origin", "*")
	s.Use(mux.CORSMethodMiddleware(r))
}

func (o *OKNO) staticHost(r *mux.Router, host string) {
	p := cfg.Path + "/sites/" + host + "/static"
	//r.StrictSlash(true).
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir( p)))
	r.StrictSlash(true).PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(p))))
	r.Headers("Access-Control-Allow-Headers", "*")
	r.Headers("Access-Control-Allow-Origin", "*")

}
