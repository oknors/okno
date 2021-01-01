package host

import (
	"github.com/gorilla/mux"
)

func djordjeMarcetinCOM() *Host {
	////////////////
	// djordje.marcetin.COM
	////////////////
	host := &Host{
		Name: "Djordje Marcetin",
		Slug: "djordje_marcetin_com",
		Host: "djordje.marcetin.com",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
