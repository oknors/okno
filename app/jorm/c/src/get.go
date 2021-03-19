package csrc


// GetCoinSources updates the available coin information sources
func GetCoinSources() {
	//go getCryptoCompare()
	// getCoinCodex()
	go getCoinGecko()

	return
}

func insertString(i, s string) string {
	if i != "" {
		if s == "" {
			s = i
		}
	}
	return s
}
func stringSlice(i []interface{}) (s []string) {
	for _, iSingle := range i {
		if iSingle != nil {
			s = append(s, iSingle.(string))
		}
	}
	return
}

func insertStringSlice(i, s []string) []string {
	for _, iSingle := range i {
		if iSingle != "" {
			if len(s) > 0 {
				for _, sSingle := range s {
					if iSingle != sSingle {
						s = append(s, iSingle)
					}
				}
			}else{
				s = append(s, iSingle)
			}
		}
	}
	return s
}

func insertInt(i, s int) int {
	if i != 0 {
		if s == 0 {
			s = i
		}
	}
	return s
}
func insertFloat(i, s float64) float64 {
	if i != 0 {
		if s == 0 {
			s = i
		}
	}
	return s
}
