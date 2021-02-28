package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/post"
	"net/http"
)

// Create appends post path prefix for a database write
func (j *JDB) CreateHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	err := r.ParseForm()
	if err != nil {
		// Handle error
	}
	var post post.Post
	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&post, r.PostForm)
	if err != nil {
		// Handle error
	}
	j.Create(path, col, id, post)
}

// Create appends post path prefix for a database write
func (j *JDB) Create(path, col, id string, post post.Post) {
	if err := j.db.Write(path+"/"+col, id, post); err != nil {
		fmt.Println("Error", err)
	}
}
