package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/jorm/h"
)

func (o *OKNO) wing(r *mux.Router) {
	////////////////
	// jorm
	////////////////
	wing := NewWingCal()

	s := r.Host("wing.marcetin.com").Subrouter()
	s.StrictSlash(true)

	s.HandleFunc("/", h.HomeHandler)

	//posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	s.HandleFunc("/radovi", wing.VrsteRadova).Methods("GET")
	s.HandleFunc("/radovi/{id}", wing.PodvrsteRadova).Methods("GET")
	s.HandleFunc("/radovi/{id}/{el}", wing.Elementi).Methods("GET")
	s.HandleFunc("/radovi/{id}/{el}/{e}", wing.Element).Methods("GET")


	s.Headers("Access-Control-Allow-Origin", "*")

}
