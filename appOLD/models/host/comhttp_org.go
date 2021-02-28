package host

import (
	"github.com/gorilla/mux"
)

func comhttpORG() *Host {
	////////////////
	// com-httpORG
	////////////////
	host := &Host{
		Name: "COM-HTTP.ORG",
		Slug: "comhttp_org",
		Host: "com-http.org",
	}
	routes := func(r *mux.Router) {
		host.testRoutes(r)
	}
	host.Routes = routes
	return host
}
