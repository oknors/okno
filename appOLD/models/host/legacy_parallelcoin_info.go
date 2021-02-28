package host

func legacyParallelcoinINFO() *Host {
	////////////////
	// legacy.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Legacy",
		Slug: "legacy_parallelcoin_info",
		Host: "legacy.parallelcoin.info",
	}
	//routes := func(r *mux.Router) {
	//	host.testRoutes(r)
	//}
	h.Routes = h.static()
	return h
}
