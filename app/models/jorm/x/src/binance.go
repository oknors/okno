package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/models/jorm/x"
)

type BinanceExchange struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                 string   `json:"symbol"`
		Status                 string   `json:"status"`
		BaseAsset              string   `json:"baseAsset"`
		BaseAssetPrecision     int      `json:"baseAssetPrecision"`
		QuoteAsset             string   `json:"quoteAsset"`
		QuotePrecision         int      `json:"quotePrecision"`
		OrderTypes             []string `json:"orderTypes"`
		IcebergAllowed         bool     `json:"icebergAllowed"`
		IsSpotTradingAllowed   bool     `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed bool     `json:"isMarginTradingAllowed"`
		Filters                []struct {
			FilterType       string `json:"filterType"`
			MinPrice         string `json:"minPrice,omitempty"`
			MaxPrice         string `json:"maxPrice,omitempty"`
			TickSize         string `json:"tickSize,omitempty"`
			MultiplierUp     string `json:"multiplierUp,omitempty"`
			MultiplierDown   string `json:"multiplierDown,omitempty"`
			AvgPriceMins     int    `json:"avgPriceMins,omitempty"`
			MinQty           string `json:"minQty,omitempty"`
			MaxQty           string `json:"maxQty,omitempty"`
			StepSize         string `json:"stepSize,omitempty"`
			MinNotional      string `json:"minNotional,omitempty"`
			ApplyToMarket    bool   `json:"applyToMarket,omitempty"`
			Limit            int    `json:"limit,omitempty"`
			MaxNumAlgoOrders int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
	} `json:"symbols"`
}
type BinanceExchangeTickers struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

func getBinanceExchange() {
	fmt.Println("GetBinanceExchangeStart")
	exchangeRaw := BinanceExchange{}
	tickersRaw := []BinanceExchangeTickers{}

	slug := "binance"
	var exchange x.Exchange
	exchange.Name = "Binance"
	exchange.Slug = slug
	resps, err := http.Get("https://api.binance.com/api/v1/exchangeInfo")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://api.binance.com/api/v1/ticker/24hr")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]BinanceExchangeTickers)

	for _, ticker := range tickersRaw {
		tickers[ticker.Symbol] = ticker
	}

	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw.Symbols {
		cur := x.Currency{
			Symbol: marketSrc.BaseAsset,
			Ask:    tickers[marketSrc.Symbol].AskPrice,
			Bid:    tickers[marketSrc.Symbol].BidPrice,
			High:   tickers[marketSrc.Symbol].HighPrice,
			Last:   tickers[marketSrc.Symbol].LastPrice,
			Low:    tickers[marketSrc.Symbol].LowPrice,
			Volume: tickers[marketSrc.Symbol].Volume,
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

	fmt.Println("GetBinanceExchangeDone")

}
