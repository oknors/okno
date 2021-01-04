package host

func logParallelcoinINFO() *Host {
	////////////////
	// log.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Log",
		Slug: "log_parallelcoin_info",
		Host: "log.parallelcoin.info",
	}
	//routes := func(r *mux.Router) {
	//	host.testRoutes(r)
	//}
	h.Routes = h.static()
	return h
}
