package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
)

type Host struct {
	Name string
	Slug string
	Host string
	//Routes func(r *mux.Router)
}

func (o *OKNO) GetHosts() map[string]Host {
	hosts := make(map[string]Host)
	err := o.Database.Read("conf", "hosts", &hosts)
	utl.ErrorLog(err)
	return hosts
}

func (h *Host) testRoutes(r *mux.Router) {
	s := h.sub(r)
	s.Host(h.Host).Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Host: %v\n", h.Name)
		})
}

func (h *Host) sub(r *mux.Router) *mux.Router {
	s := r.Host(h.Host).Subrouter()
	return s
}
