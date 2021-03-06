package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/mod"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID      string `schema:"id,required"`
	Title   string
	Slug    string
	Date    time.Time
	Excerpt string
	Active  bool
	Order   int
}
type Posts []Post

func (p Posts) Len() int           { return len(p) }
func (p Posts) Less(i, j int) bool { return p[i].Order < p[j].Order }
func (p Posts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (o *OKNO) api(r *mux.Router) {
	////////////////
	// api
	////////////////
	a := r.Host("api.okno.rs").Subrouter()
	a.StrictSlash(true)

	a.HandleFunc("/", homeHandler)
	a.HandleFunc("/{site}/{col}/{id}", o.viewPost).Methods("GET")
	a.HandleFunc("/{site}/{col}", o.viewAllPosts).Methods("GET")
	a.HandleFunc("/{site}/{col}/{per}/{page}/{truncate}", o.viewPosts).Methods("GET")
	a.Headers("Access-Control-Allow-Origin", "*")
}

// HomeHandler handles a request for (?)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The Truth Is Out There..."))
}

func (o *OKNO) viewPosts(w http.ResponseWriter, r *http.Request) {
	per, _ := strconv.Atoi(mux.Vars(r)["per"])
	page, _ := strconv.Atoi(mux.Vars(r)["page"])
	trn, _ := strconv.Atoi(mux.Vars(r)["truncate"])
	site := mux.Vars(r)["site"]
	col := mux.Vars(r)["col"]

	posts := Posts{}

	postsRaw, err := o.Database.ReadAll("/sites/" + site + "/jdb/" + col)
	utl.ErrorLog(err)
	for _, postInterface := range postsRaw {
		var rawPost mod.Post
		err := json.Unmarshal([]byte(postInterface), &rawPost)
		utl.ErrorLog(err)
		if !rawPost.IsDraft {
			p := Post{
				ID:      rawPost.ID,
				Title:   rawPost.Title,
				Slug:    rawPost.Slug,
				Date:    rawPost.CreatedAt,
				Order:   rawPost.Order,
				Excerpt: truncate(rawPost.Content, trn),
			}
			posts = append(posts, p)
		}
	}
	sort.Sort(posts)

	pn := len(posts)
	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   int(pn) / per,
		"posts":       posts,
		"postsNumber": pn,
	}
	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (o *OKNO) viewAllPosts(w http.ResponseWriter, r *http.Request) {
	site := mux.Vars(r)["site"]
	col := mux.Vars(r)["col"]

	posts := mod.Posts{}

	postsRaw, err := o.Database.ReadAll("/sites/" + site + "/jdb/" + col)
	utl.ErrorLog(err)
	for _, postInterface := range postsRaw {
		var rawPost mod.Post
		err := json.Unmarshal([]byte(postInterface), &rawPost)
		utl.ErrorLog(err)
		if !rawPost.IsDraft {
			posts = append(posts, rawPost)
		}
	}
	sort.Sort(posts)

	out, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (o *OKNO) viewPost(w http.ResponseWriter, r *http.Request) {
	site := mux.Vars(r)["site"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["id"]
	post := o.Database.ReadPost("/sites/"+site+"/jdb", col, id)
	out, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
