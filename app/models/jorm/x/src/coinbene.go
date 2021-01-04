package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/models/jorm/x"
)

type CoinBeneExchange struct {
	Symbol []struct {
		Ticker      string `json:"ticker"`
		BaseAsset   string `json:"baseAsset"`
		QuoteAsset  string `json:"quoteAsset"`
		TakerFee    string `json:"takerFee"`
		MakerFee    string `json:"makerFee"`
		TickSize    string `json:"tickSize"`
		LotStepSize string `json:"lotStepSize"`
		MinQuantity string `json:"minQuantity"`
	} `json:"symbol"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type CoinBeneExchangeTickers struct {
	Ticker    []CoinBeneExchangeTicker `json:"ticker"`
	Status    string                   `json:"status"`
	Timestamp int64                    `json:"timestamp"`
}

type CoinBeneExchangeTicker struct {
	Symbol     string `json:"symbol"`
	Two4HrHigh string `json:"24hrHigh"`
	Last       string `json:"last"`
	Two4HrVol  string `json:"24hrVol"`
	Ask        string `json:"ask"`
	Two4HrLow  string `json:"24hrLow"`
	Bid        string `json:"bid"`
	Two4HrAmt  string `json:"24hrAmt"`
}

func getCoinBeneExchange() {
	fmt.Println("GetCoinBeneExchangeStart")
	exchangeRaw := CoinBeneExchange{}
	tickersRaw := CoinBeneExchangeTickers{}

	slug := "coinbene"
	var exchange x.Exchange
	exchange.Name = "CoinBene"
	exchange.Slug = slug
	resps, err := http.Get("http://api.coinbene.com/v1/market/symbol")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("http://api.coinbene.com/v1/market/ticker?symbol=all")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]CoinBeneExchangeTicker)
	for _, ticker := range tickersRaw.Ticker {
		tickers[ticker.Symbol] = ticker
	}
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw.Symbol {
		cur := x.Currency{
			Symbol: marketSrc.BaseAsset,
			Ask:    tickers[marketSrc.Ticker].Ask,
			Bid:    tickers[marketSrc.Ticker].Bid,
			High:   tickers[marketSrc.Ticker].Two4HrHigh,
			Last:   tickers[marketSrc.Ticker].Last,
			Low:    tickers[marketSrc.Ticker].Two4HrLow,
			Volume: tickers[marketSrc.Ticker].Two4HrVol,
		}
		_, ok := markets[marketSrc.QuoteAsset]
		if !ok {
			markets[marketSrc.QuoteAsset] = []x.Currency{}
		}
		markets[marketSrc.QuoteAsset] = append(markets[marketSrc.QuoteAsset], cur)
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
	jdb.WriteExchange(slug, exchange)

	fmt.Println("GetCoinBeneExchangeDone")

}
