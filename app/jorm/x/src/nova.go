package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/jorm/x"
)

type NovaExchange struct {
	Markets []struct {
		Ask          string `json:"ask"`
		Basecurrency string `json:"basecurrency"`
		Bid          string `json:"bid"`
		Change24H    string `json:"change24h"`
		Currency     string `json:"currency"`
		Disabled     int    `json:"disabled"`
		High24H      string `json:"high24h"`
		LastPrice    string `json:"last_price"`
		Low24H       string `json:"low24h"`
		Marketid     int    `json:"marketid"`
		Marketname   string `json:"marketname"`
		Volume24H    string `json:"volume24h"`
	} `json:"markets"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func getNovaExchange() {
	fmt.Println("GetNovaExchangeStart")
	marketsRaw := NovaExchange{}
	slug := "nova"
	var exchange x.Exchange
	exchange.Name = "Nova"
	exchange.Slug = slug
	respcs, err := http.Get("https://novaexchange.com/remote/v2/markets/")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)

	var cursor x.Market
	var beforeBase string
	for _, marketSrc := range marketsRaw.Markets {
		if beforeBase != marketSrc.Basecurrency {
			if cursor.Symbol != "" {
				exchange.Markets = append(exchange.Markets, cursor)
			}
			cursor = x.Market{Symbol: marketSrc.Basecurrency}
		}
		cursor.Currencies = append(cursor.Currencies,
			x.Currency{
				Symbol: marketSrc.Currency,
				Ask:    marketSrc.Ask,
				Bid:    marketSrc.Bid,
				High:   marketSrc.High24H,
				Last:   marketSrc.LastPrice,
				Low:    marketSrc.Low24H,
				Volume: marketSrc.Volume24H,
			})
		beforeBase = marketSrc.Basecurrency
	}

	jdb.WriteExchange(slug, exchange)

	fmt.Println("GetNovaExchangeDone")

}
