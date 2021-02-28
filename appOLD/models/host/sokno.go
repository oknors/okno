package host

import (
	"github.com/gorilla/mux"
	"net/http"
)

func sokno() *Host {
	////////////////
	// s.okno.RS
	////////////////
	h := &Host{
		Name: "sokno",
		Slug: "sokno",
		Host: "s.okno.rs",
	}
	routes := func(r *mux.Router) {
		s := h.sub(r)
		s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	}
	h.Routes = routes
	return h
}
