package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p9c/okno/app/models/post"
	"net/http"
)

// Create appends post path prefix for a database write
func (j *JDB) Create(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	data := post.Post{}
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

	if err := j.db.Write(path+"/"+col, id, data); err != nil {
		fmt.Println("Error", err)
	}
}
