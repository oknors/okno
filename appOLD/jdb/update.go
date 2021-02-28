package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/post"
	"net/http"
)

// Update appends post path prefix for a database write
func (j *JDB) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	data := post.Post{}
	if err := j.db.Write(path+"/"+col, id, data); err != nil {
		fmt.Println("Error", err)
	}
	return
}

// Update appends post path prefix for a database write
func (j *JDB) Update(path, col, id string) {
	data := post.Post{}
	if err := j.db.Write(path+"/"+col, id, data); err != nil {
		fmt.Println("Error", err)
	}
	return
}
