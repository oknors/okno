package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Delete  data from the database
func (j *JDB) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	j.DeletePost(path, col, id)
	return
}

// Delete  data from the database
func (j *JDB) DeletePost(path, col, id string) {
	if err := j.Delete(path+"/"+col, id); err != nil {
		fmt.Println("Error", err)
	}
	return
}
