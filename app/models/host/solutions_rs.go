package host

import (
	"github.com/gorilla/mux"
)

func solutionsRS() *Host {
	////////////////
	// solutions.RS
	////////////////
	host := &Host{
		Name: "W-ing Solutions",
		Slug: "solutions_rs",
		Host: "solutions.rs",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
