package jdb

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Delete  data from the database
func (j *JDB) Delete(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["host"]
	col := mux.Vars(r)["col"]
	id := mux.Vars(r)["slug"]
	if err := j.db.Delete(path+"/"+col, id); err != nil {
		fmt.Println("Error", err)
	}
	return
}
