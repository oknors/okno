package jdb

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/mod"
	"net/http"
)

// Update appends post path prefix for a database write
func (j *JDB) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	j.UpdatePost(path, col, id, mod.Post{})
	return
}

// Update appends post path prefix for a database write
func (j *JDB) UpdatePost(path, col, id string, post mod.Post) error {
	return j.Write(path+"/"+col, id, post)
}
