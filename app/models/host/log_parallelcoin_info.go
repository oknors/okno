package host

import (
	"github.com/gorilla/mux"
)

func logParallelcoinINFO() *Host {
	////////////////
	// log.parallelcoin.INFO
	////////////////
	host := &Host{
		Name: "ParallelCoin Log",
		Slug: "log_parallelcoin_info",
		Host: "log.parallelcoin.info",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
