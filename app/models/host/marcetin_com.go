package host

import (
	"github.com/gorilla/mux"
	"net/http"
)

func marcetinCOM() *Host {
	////////////////
	// marcetin.COM
	////////////////
	host := &Host{
		Name: "Marcetin",
		Slug: "marcetin_com",
		Host: "marcetin.com",
	}
	routes := func(r *mux.Router) {
		s := r.Host(host.Host).Subrouter()
		s.StrictSlash(true)
		//s.Headers("X-Requested-With", "XMLHttpRequest")
		s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/public/parallelcoininfo")))
	}
	host.Routes = routes
	return host
}
