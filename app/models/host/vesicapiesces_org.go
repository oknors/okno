package host

import (
	"github.com/gorilla/mux"
)

func vesicaPiescesORG() *Host {
	////////////////
	// VesicaPiescesORG
	////////////////
	host := &Host{
		Name: "Vesica Piesces",
		Slug: "vesicapiesces_org",
		Host: "vesicapiesces.org",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
