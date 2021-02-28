package host

func downloadParallelcoinINFO() *Host {
	////////////////
	// download.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Download",
		Slug: "download_parallelcoin_info",
		Host: "download.parallelcoin.info",
	}
	h.Routes = h.static()
	return h
}
