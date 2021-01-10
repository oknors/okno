package host

func whitepaperParallelcoinINFO() *Host {
	////////////////
	// whitepaper.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Whitepaper",
		Slug: "whitepaper_parallelcoin_info",
		Host: "whitepaper.parallelcoin.info",
	}
	h.Routes = h.static()
	return h
}
