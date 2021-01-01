package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p9c/okno/app/models/post"
	"net/http"
)

var (
	out interface{}
)

// Renders okno posts JSON
func (hs *Handlers) GetJSON(w http.ResponseWriter, r *http.Request) {
	hs.Host(mux.Vars(r)["host"])
	col := mux.Vars(r)["col"]
	slug := mux.Vars(r)["slug"]
	hs.Collection(col)
	if slug != "" {
		out = hs.getPost(slug)
	} else {
		switch col {
		case "pages":
			posts := make(map[string]interface{})
			for _, p := range *hs.getPosts() {
				posts[p.Slug] = p
			}
			out = posts
		default:
			out = hs.getPosts()
		}
	}
	render(w, out)
}

// Renders JSON
func render(w http.ResponseWriter, p interface{}) {
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (hs *Handlers) getPosts() *[]post.Post {
	var posts []post.Post
	ps, _ := hs.ReadAll()
	for _, postInterface := range ps {
		var p post.Post
		if err := json.Unmarshal([]byte(postInterface), &p); err != nil {
			fmt.Println("Error", err)
		}
		posts = append(posts, p)
	}
	return &posts
}

func (hs *Handlers) getPost(slug string) *post.Post {
	p := post.Post{}
	err := hs.Read(slug, &p)
	if err != nil {
	}
	return &p
}
