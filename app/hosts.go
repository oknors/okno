package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
)

type Host struct {
	Name     string
	Slug     string
	Host     string
	Template string
	Subs     map[string]Sub
	//Routes func(r *mux.Router)
}
type Sub struct {
	Name     string
	Slug     string
	Template string
}

func (o *OKNO) GetHosts() map[string]Host {
	hosts := make(map[string]Host)
	err := jdb.JDB.Read("conf", "hosts", &hosts)
	utl.ErrorLog(err)
	return hosts
}

func (h *Host) testRoutes(r *mux.Router) {
	s := h.domain(r)
	s.Host(h.Host).Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Host: %v\n", h.Name)
		})
}

func (h *Host) domain(r *mux.Router) *mux.Router {
	s := r.Host(h.Host).Subrouter()
	return s
}
func (h *Host) sub(r *mux.Router) *mux.Router {
	s := r.Host("{site}." + h.Host).Subrouter()
	return s
}
