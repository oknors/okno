package host

import (
	"github.com/gorilla/mux"
)

func okno() *Host {
	////////////////
	// okno.RS
	////////////////
	h := &Host{
		Name: "okno",
		Slug: "okno",
		Host: "okno.rs",
	}

	//hs := handlers.Handlers{jdb.NewJDB(db, h.Slug)}

	routes := func(r *mux.Router) {
		s := h.sub(r)
		s.Headers("X-Requested-With", "XMLHttpRequest")
		//s.Path("/").HandlerFunc(hs.Homepage()).Name("okno")
		//s.Path("/{coin}").HandlerFunc(hs.Homepage()).Name("comhttpus")
		//r.Host(okno).Path("/coins/").HandlerFunc(rts.Coins).Methods("GET")
		//r.Host(okno).Path("/bitnodes/").HandlerFunc(rts.BitNodes).Methods("GET")
		//r.Host(okno).Path("/bitnodes/{coin}/{nodeid}").HandlerFunc(rts.BitNode).Methods("GET")
		//
		//r.Host("{coin}." + okno).Path("/").HandlerFunc(rts.Coin).Methods("GET")
		//r.Host("{coin}." + okno).Path("/favicon.ico").HandlerFunc(rts.Ico).Name("ico")
		//
		//r.Host("api." + okno).Path("/{coin}/i").HandlerFunc(rts.ApiInfo).Name("info")
	}
	h.Routes = routes

	//seed.SeedOKNO(jdb.NewJDB(db, h.Slug))
	return h
}
