package host

import (
	"github.com/gorilla/mux"
	"net/http"
)

func parallelcoinINFO() *Host {
	////////////////
	// parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin ",
		Slug: "parallelcoin_info",
		Host: "parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		s := r.Host(h.Host).Subrouter()
		s.StrictSlash(true)
		//s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/parallelcoin/__sapper__/export")))
	}
	h.Routes = routes
	//seed.SeedParallelCoin(jdb.NewJDB(db, h.Slug))
	return h
}
