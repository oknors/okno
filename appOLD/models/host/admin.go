package host

func admin() *Host {
	////////////////
	// admin.okno.rs
	////////////////
	h := &Host{
		Name: "Admin",
		Slug: "admin",
		Host: "admin.okno.rs",
	}
	//h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}
	//routes := func(r *mux.Router) {
	//	s := h.sub(r)
	//	s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/public/admin")))
	//}
	h.Routes = h.static()
	return h
}
