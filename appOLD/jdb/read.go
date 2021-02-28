package jdb

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/post"
	"net/http"
)

// // Response Handler
func (j *JDB) ReadHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	render(w, j.Read(path, col, id))
}

// // ReadData appends 'data' path prefix for a database read
func (j *JDB) Read(path, col, id string) post.Post {
	data := post.Post{}
	err := j.db.Read(path+"/"+col, id, &data)
	if err != nil {
	}
	return data
}

// Rresponse Handler.
func (j *JDB) ReadCollectionHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	render(w, j.ReadCollection(path, col))
}

// Read all items from the database, unmarshaling the response.
func (j *JDB) ReadCollection(path, col string) (out interface{}) {
	var posts []map[string]string
	data, err := j.db.ReadAll(path + "/" + col)
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, postInterface := range data {
		var p post.Post
		if err := json.Unmarshal([]byte(postInterface), &p); err != nil {
			fmt.Println("Error", err)
		}
		posts = append(posts, map[string]string{
			"Title": p.Title,
			"Slug":  p.Slug,
			"Date":  p.CreatedAt.String(),
		})
	}
	switch col {
	case "pages":
		pages := make(map[string]interface{})
		for _, p := range posts {
			pages[p["Slug"]] = p
		}
		out = pages
	default:
		out = posts
	}
	return
}
