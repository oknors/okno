package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oknors/okno/appOLD/models/jorm/jdb"
	"github.com/oknors/okno/appOLD/models/jorm/x"
)

type HuobiExchange struct {
	Status string `json:"status"`
	Data   []struct {
		BaseCurrency    string  `json:"base-currency"`
		QuoteCurrency   string  `json:"quote-currency"`
		PricePrecision  int     `json:"price-precision"`
		AmountPrecision int     `json:"amount-precision"`
		SymbolPartition string  `json:"symbol-partition"`
		Symbol          string  `json:"symbol"`
		State           string  `json:"state"`
		ValuePrecision  int     `json:"value-precision"`
		MinOrderAmt     int     `json:"min-order-amt"`
		MaxOrderAmt     int     `json:"max-order-amt"`
		MinOrderValue   float64 `json:"min-order-value"`
		LeverageRatio   int     `json:"leverage-ratio,omitempty"`
	} `json:"data"`
}
type HuobiExchangeTickers struct {
	Status string                  `json:"status"`
	Ts     int64                   `json:"ts"`
	Data   []HuobiExchangeCurrency `json:"data"`
}
type HuobiExchangeCurrency struct {
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
	Vol    float64 `json:"vol"`
	Symbol string  `json:"symbol"`
}

func getHuobiExchange() {
	fmt.Println("GetHuobiExchangeStart")
	exchangeRaw := HuobiExchange{}
	tickersRaw := HuobiExchangeTickers{}

	slug := "huobi"
	var exchange x.Exchange
	exchange.Name = "Huobi"
	exchange.Slug = slug
	resps, err := http.Get("https://api.huobi.pro/v1/common/symbols")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://api.huobi.pro/market/tickers")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]HuobiExchangeCurrency)

	for _, ticker := range tickersRaw.Data {
		tickers[ticker.Symbol] = ticker
	}
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw.Data {
		cur := x.Currency{
			Symbol: strings.ToUpper(marketSrc.BaseCurrency),
			// Ask:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Ask),
			// Bid:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Bid),
			High: fmt.Sprintf("%f", tickers[marketSrc.Symbol].High),
			// Last:   fmt.Sprintf("%f", tickers[marketSrc.Symbol].Last),
			Low:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Low),
			Volume: fmt.Sprintf("%f", tickers[marketSrc.Symbol].Vol),
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
			Symbol:     strings.ToUpper(i),
			Currencies: markets[i],
		}
		marketsSlice = append(marketsSlice, newSlice)
	}
	exchange.Markets = marketsSlice
	jdb.WriteExchange(slug, exchange)

	fmt.Println("GetHuobiExchangeDone")

}
