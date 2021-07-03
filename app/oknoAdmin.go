package app

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/mod"
	"github.com/oknors/okno/app/tpl"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"sort"
	"time"
)

var decoder = schema.NewDecoder()

func (o *OKNO) oknoAdmin(r *mux.Router) {
	okno := r.Host("admin.okno.rs").Subrouter()
	okno.HandleFunc("/", o.indexAdmin).Methods("GET")
	okno.HandleFunc("/hosts", o.readHosts).Methods("GET")
	//okno.HandleFunc("/coins", o.viewCoins).Methods("GET")
	//okno.HandleFunc("/coins/all", o.readCoins).Methods("GET")
	okno.HandleFunc("/{site}/{col}/list", o.listAdmin).Methods("GET")
	okno.HandleFunc("/{site}/config}", o.configAdmin).Methods("GET")
	okno.HandleFunc("/{site}/create", o.createAdmin).Methods("GET")
	okno.HandleFunc("/{site}/{col}/edit/{slug}", o.editAdmin).Methods("GET")
	okno.HandleFunc("/{site}/{col}/write", o.writeAdmin).Methods("POST")
	okno.StrictSlash(true).PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Path+"/admin/static"))))
}

// HomeHandler handles a request for (?)
func (o *OKNO) indexAdmin(w http.ResponseWriter, r *http.Request) {
	o.GetParallelCoin()

	data := &PageData{
		Host:  o.Hosts["okno_rs"],
		Hosts: o.hosts(),
	}
	fmt.Println("Index")
	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "index_gohtml", data)
}

// HomeHandler handles a request for (?)
func (o *OKNO) listAdmin(w http.ResponseWriter, r *http.Request) {
	var posts []mod.Post
	postsRaw, err := jdb.JDB.ReadAll("sites/" + mux.Vars(r)["site"] + "/jdb/" + mux.Vars(r)["col"])
	utl.ErrorLog(err)
	for _, postInterface := range postsRaw {
		var p mod.Post
		err := json.Unmarshal([]byte(postInterface), &p)
		utl.ErrorLog(err)
		posts = append(posts, p)
	}

	fmt.Println("o.Hosts", o.Hosts)
	fmt.Println("o.Hosts", o.Hosts[mux.Vars(r)["site"]])
	fmt.Println("o.Site", mux.Vars(r)["site"])
	data := &PageData{
		Host:  o.Hosts[mux.Vars(r)["site"]],
		Hosts: o.hosts(),
		Posts: posts,
	}
	fmt.Println("List posts", mux.Vars(r)["site"])
	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "list_gohtml", data)
}

// HomeHandler handles a request for (?)
func (o *OKNO) createAdmin(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Host:  o.Hosts[mux.Vars(r)["site"]],
		Hosts: o.hosts(),
	}
	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "editpost_gohtml", data)
}

// HomeHandler handles a request for (?)
func (o *OKNO) configAdmin(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Host:  o.Hosts[mux.Vars(r)["site"]],
		Hosts: o.hosts(),
	}
	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "config_gohtml", data)
}

func (o *OKNO) hosts() (hosts []Host) {
	for _, host := range o.Hosts {
		hosts = append(hosts, host)
	}
	return
}

func (o *OKNO) posts(path string) (posts mod.Posts) {
	postsRaw, err := jdb.JDB.ReadAll(path)
	utl.ErrorLog(err)
	for _, postInterface := range postsRaw {
		var p mod.Post
		err := json.Unmarshal([]byte(postInterface), &p)
		utl.ErrorLog(err)
		if !p.IsDraft {
			posts = append(posts, p)
		}
	}
	sort.Sort(posts)
	return
}

// HomeHandler handles a request for (?)
func (o *OKNO) editAdmin(w http.ResponseWriter, r *http.Request) {
	post := mod.Post{}
	err := jdb.JDB.Read("sites/"+mux.Vars(r)["site"]+"/jdb/"+mux.Vars(r)["col"], mux.Vars(r)["slug"], &post)
	utl.ErrorLog(err)

	data := &PageData{
		Host:  o.Hosts[mux.Vars(r)["site"]],
		Hosts: o.hosts(),
		Post:  post,
	}
	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "editpost_gohtml", data)
	fmt.Println("Read", post.Title)
}

// HomeHandler handles a request for (?)
func (o *OKNO) writeAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var post mod.Post
	decoder.Decode(&post)
	if post.ID == "" {
		post.ID = uuid.Must(uuid.NewV4()).String()
	}
	if post.CreatedAt.String() == "" {
		post.CreatedAt = time.Now()
	}
	post.UpdatedAt = time.Now()
	if post.Slug == "" {
		post.Slug = utl.MakeSlug(post.Title)
	}

	utl.ErrorLog(jdb.JDB.Write("/sites/"+mux.Vars(r)["site"]+"/jdb/"+mux.Vars(r)["col"], post.Slug, post))
	fmt.Println("Write", post.Title)
	fmt.Println("Write", post.Order)
}

func (o *OKNO) readHosts(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(o.Hosts)
	if err != nil {
	}
	fmt.Println("Hosts", o.Hosts)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
//
//func (o *OKNO) viewCoins(w http.ResponseWriter, r *http.Request) {
//	data := map[string]interface{}{
//		"coins": coin.LoadCoinsBase(false, false),
//	}
//	tpl.TemplateHandler(cfg.Path+"/admin").ExecuteTemplate(w, "coins_gohtml", data)
//}
//
//func (o *OKNO) readCoins(w http.ResponseWriter, r *http.Request) {
//	js, err := json.Marshal(coin.LoadCoinsBase(false, false))
//	utl.ErrorLog(err)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}
