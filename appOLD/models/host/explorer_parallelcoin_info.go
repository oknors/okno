package host

func explorerParallelcoinINFO() *Host {
	////////////////
	// explorer.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin explorer",
		Slug: "explorer_parallelcoin_info",
		Host: "explorer.parallelcoin.info",
	}
	h.Routes = h.static()
	return h
}
