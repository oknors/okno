package host

func parallelcoinINFO() *Host {
	////////////////
	// parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin ",
		Slug: "parallelcoin_info",
		Host: "new.parallelcoin.info",
	}
	//routes := func(r *mux.Router) {
	//	s := r.Host(h.Host).Subrouter()
	//	s.StrictSlash(true)
	//	s.Headers("X-Requested-With", "XMLHttpRequest")
	//s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/parallelcoin/__sapper__/export")))
	//}
	h.Routes = h.static()
	return h
}