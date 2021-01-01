package host

import (
	"github.com/gorilla/mux"
)

func parallelcoinIO() *Host {
	////////////////
	// parallelcoin.IO
	////////////////
	host := &Host{
		Name: "ParallelCoin",
		Slug: "parallelcoin",
		Host: "parallelcoin.io",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
