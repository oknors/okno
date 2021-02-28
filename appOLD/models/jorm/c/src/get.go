package csrc

// GetCoinSources updates the available coin information sources
func GetCoinSources() {
	go getCryptoCompare()
	// getCoinCodex()
	// getCoinGecko()

	return
}

func insertData(i, s string) string {
	if i != "" {
		if s == "" {
			s = i
		}
	}
	return s
}
