package app

func (o *OKNO) GetParallelCoin() {
	o.Hosts["parallelcoin_info"] = Host{
		Name: "ParallelCoin",
		Slug: "parallelcoin_info",
		Host: "parallelcoin.info",
	}
	o.Hosts["credits_parallelcoin_info"] = Host{
		Name: "ParallelCoin Credits",
		Slug: "credits_parallelcoin_info",
		Host: "credits.parallelcoin.info",
	}
	o.Hosts["dev_parallelcoin_info"] = Host{
		Name: "ParallelCoin Development",
		Slug: "dev_parallelcoin_info",
		Host: "dev.parallelcoin.info",
	}
	o.Hosts["log_parallelcoin_info"] = Host{
		Name: "ParallelCoin's Starlog",
		Slug: "log_parallelcoin_info",
		Host: "log.parallelcoin.info",
	}
	o.Hosts["faq_parallelcoin_info"] = Host{
		Name: "ParallelCoin's frequently asked questions",
		Slug: "faq_parallelcoin_info",
		Host: "faq.parallelcoin.info",
	}
	o.Hosts["links_parallelcoin_info"] = Host{
		Name: "ParallelCoin Links",
		Slug: "links_parallelcoin_info",
		Host: "links.parallelcoin.info",
	}
	o.Hosts["specs_parallelcoin_info"] = Host{
		Name: "ParallelCoin Specifications",
		Slug: "specs_parallelcoin_info",
		Host: "specs.parallelcoin.info",
	}

	o.Hosts["spore_parallelcoin_info"] = Host{
		Name: "Spore protocol",
		Slug: "spore_parallelcoin_info",
		Host: "spore.parallelcoin.info",
	}

	o.Hosts["whitepaper_parallelcoin_info"] = Host{
		Name: "ParallelCoin Whitepaper",
		Slug: "whitepaper_parallelcoin_info",
		Host: "whitepaper.parallelcoin.info",
	}
	o.Hosts["wiki_parallelcoin_info"] = Host{
		Name: "ParallelCoin Wiki",
		Slug: "wiki_parallelcoin_info",
		Host: "wiki.parallelcoin.info",
	}
	return
}