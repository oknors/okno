package host

import (
	"github.com/gorilla/mux"
	"net/http"
)

func docParallelcoinINFO() *Host {
	////////////////
	// doc.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Documentation",
		Slug: "doc_parallelcoin_info",
		Host: "doc.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		s := r.Host(h.Host).Subrouter()
		s.StrictSlash(true)
		//s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/doc/public")))
	}
	h.Routes = routes
	//seed.SeedParallelCoin(jdb.NewJDB(db, h.Slug))
	return h
}
