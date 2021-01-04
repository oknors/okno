package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/models/jorm/x"
)

type BitMartExchange struct {
	ID                string `json:"id"`
	BaseCurrency      string `json:"base_currency"`
	QuoteCurrency     string `json:"quote_currency"`
	QuoteIncrement    string `json:"quote_increment"`
	BaseMinSize       string `json:"base_min_size"`
	BaseMaxSize       string `json:"base_max_size"`
	PriceMinPrecision int    `json:"price_min_precision"`
	PriceMaxPrecision int    `json:"price_max_precision"`
	Expiration        string `json:"expiration"`
}
type BitMartExchangeTickers struct {
	Volume       string `json:"volume"`
	Ask1         string `json:"ask_1"`
	BaseVolume   string `json:"base_volume"`
	LowestPrice  string `json:"lowest_price"`
	Bid1         string `json:"bid_1"`
	HighestPrice string `json:"highest_price"`
	Ask1Amount   string `json:"ask_1_amount"`
	CurrentPrice string `json:"current_price"`
	Fluctuation  string `json:"fluctuation"`
	SymbolID     string `json:"symbol_id"`
	URL          string `json:"url"`
	Bid1Amount   string `json:"bid_1_amount"`
}

func getBitMartExchange() {
	fmt.Println("GetBitMartExchangeStart")
	exchangeRaw := []BitMartExchange{}
	tickersRaw := []BitMartExchangeTickers{}

	slug := "bitmart"
	var exchange x.Exchange
	exchange.Name = "BitMart"
	exchange.Slug = slug
	resps, err := http.Get("https://openapi.bitmart.com/v2/symbols_details")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://openapi.bitmart.com/v2/ticker")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]BitMartExchangeTickers)

	for _, ticker := range tickersRaw {
		tickers[ticker.SymbolID] = ticker
	}

	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw {
		cur := x.Currency{
			Symbol: marketSrc.BaseCurrency,
			Ask:    tickers[marketSrc.ID].Ask1,
			Bid:    tickers[marketSrc.ID].Bid1,
			High:   tickers[marketSrc.ID].HighestPrice,
			Last:   tickers[marketSrc.ID].CurrentPrice,
			Low:    tickers[marketSrc.ID].LowestPrice,
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
	jdb.WriteExchange(slug, exchange)

	fmt.Println("GetBitMartExchangeDone")

}
