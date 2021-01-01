package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p9c/okno/app/models/post"
	"net/http"
)

// Update appends post path prefix for a database write
func (j *JDB) Update(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	data := post.Post{}
	if err := j.db.Write(path+"/"+col, id, data); err != nil {
		fmt.Println("Error", err)
	}
	return
}
