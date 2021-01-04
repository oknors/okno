package xsrc

// GetCoinSources updates the available coin information sources
func GetExchangeSources() {
	go getNovaExchange()
	go getIDAXExchange()
	go getKuCoinExchange()
	go getLBankExchange()
	go getPoloniexExchange()
	go getDigiFinexExchange()
	go getAltillyExchange()
	go getBitTrexExchange()
	go getBinanceExchange()
	go getCoinBeneExchange()
	go getHuobiExchange()
	go getBitMartExchange()
	go getHitBTCExchange()
	go getBitZExchange()
	return
}
