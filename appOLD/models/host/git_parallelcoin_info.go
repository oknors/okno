package host

func gitParallelcoinINFO() *Host {
	////////////////
	// git.parallelcoin.IO
	////////////////
	h := &Host{
		Name: "ParallelCoin GitHub Status",
		Slug: "git_parallelcoin_info",
		Host: "git.parallelcoin.info",
	}
	//routes := func(r *mux.Router) {
	//	s := r.Host(host.Host).Subrouter()
	//	s.StrictSlash(true)
	//s.Headers("X-Requested-With", "XMLHttpRequest")
	//s.PathPrefix("/").Handler(http.FileServer(http.Dir("js/git/public")))
	//}
	h.Routes = h.static()
	return h
}
