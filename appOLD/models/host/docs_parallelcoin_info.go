package host

func docsParallelcoinINFO() *Host {
	////////////////
	// doc.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Documentation",
		Slug: "docs_parallelcoin_info",
		Host: "docs.parallelcoin.info",
	}
	h.Routes = h.static()
	return h
}
