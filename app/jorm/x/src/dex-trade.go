package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/jorm/x"
)

func getDexTradeExchange() {
	fmt.Println("Get Dex Trade Exchange Start")

	exchangeRaw := make(map[string]interface{})
	tickersRaw := make(map[string]interface{})


	slug := "dex-trade"
	var exchange x.Exchange
	exchange.Name = "Dex Trade"
	exchange.Slug = slug
	resps, err := http.Get("https://api.dex-trade.com/v1/public/symbols")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://api.altilly.com/api/public/ticker")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]interface{})

	for _, ticker := range tickersRaw {
		tickers[ticker.Symbol] = ticker
	}

	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw {
		cur := x.Currency{
			Symbol: marketSrc.BaseCurrency,
			Ask:    tickers[marketSrc.ID].Ask,
			Bid:    tickers[marketSrc.ID].Bid,
			High:   tickers[marketSrc.ID].High,
			Last:   tickers[marketSrc.ID].Last,
			Low:    tickers[marketSrc.ID].Low,
			Volume: tickers[marketSrc.ID].Volume,
		}
		_, ok := markets[marketSrc.QuoteCurrency]
		if !ok {
			markets[marketSrc.QuoteCurrency] = []x.Currency{}
		}
		markets[marketSrc.QuoteCurrency] = append(markets[marketSrc.QuoteCurrency], cur)
	}
	var marketsSlice x.Markets
	for i := range markets {
		newSlice := x.Market{
			Symbol:     i,
			Currencies: markets[i],
		}
		marketsSlice = append(marketsSlice, newSlice)
	}
	exchange.Markets = marketsSlice
	//jdb.WriteExchange(slug, exchange)

	fmt.Println("Get Dex Trade Exchange Done")

}
