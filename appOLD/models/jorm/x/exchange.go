package x

import (
	"encoding/json"
	"fmt"

	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/appOLD/models/jorm/jdb"
)

type Exchange struct {
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Url     string  `json:"url"`
	Logo    string  `json:"logo"`
	Markets Markets `json:"markets"`
}
type Currency struct {
	Symbol string `json:"symbol"`
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	High   string `json:"high"`
	Last   string `json:"last"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
}

type Exchanges []Exchange

type Market struct {
	Symbol     string     `json:"symbol"`
	Currencies []Currency `json:"currencies"`
}
type Markets []Market

type CoinMarket struct {
	Exchange     string   `json:"exchange"`
	ExchangeSlug string   `json:"exslug"`
	Market       string   `json:"market"`
	Ticker       Currency `json:"ticker"`
}
type CoinMarkets []CoinMarket

// ReadAllExchanges reads in all of the data about all coins in the database
func ReadAllExchanges() Exchanges {
	exchanges := jdb.ReadData(cfg.Web + "/exchanges")
	ex := make(Exchanges, len(exchanges))
	for i := range exchanges {
		if err := json.Unmarshal(exchanges[i], &ex[i]); err != nil {
			fmt.Println("Error", err)
		}
	}
	jdb.DB.Write(cfg.Web, "exchanges", ex)

	return ex
}
