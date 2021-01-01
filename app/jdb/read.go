package jdb

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p9c/okno/app/models/post"
	"net/http"
)

// // ReadData appends 'data' path prefix for a database read
func (j *JDB) Read(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	data := post.Post{}
	err := j.db.Read(path+"/"+col, id, &data)
	if err != nil {
	}
	render(w, data)
}

// Read all items from the database, unmarshaling the response.
func (j *JDB) ReadCollection(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]

	var posts []post.Post
	data, err := j.db.ReadAll(path + "/" + col)
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, postInterface := range data {
		var p post.Post
		if err := json.Unmarshal([]byte(postInterface), &p); err != nil {
			fmt.Println("Error", err)
		}
		posts = append(posts, p)
	}
	var out interface{}
	switch col {
	case "pages":
		pages := make(map[string]interface{})
		for _, p := range posts {
			pages[p.Slug] = p
		}
		out = pages
	default:
		out = posts
	}
	render(w, out)
}
