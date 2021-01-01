package host

import (
	"github.com/gorilla/mux"
	"net/http"
)

func gitParallelcoinINFO() *Host {
	////////////////
	// git.parallelcoin.IO
	////////////////
	host := &Host{
		Name: "ParallelCoin GitHub Status",
		Slug: "git_parallelcoin_info",
		Host: "git.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		s := r.Host(host.Host).Subrouter()
		s.StrictSlash(true)
		//s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/git/public")))
	}
	host.Routes = routes
	return host
}
