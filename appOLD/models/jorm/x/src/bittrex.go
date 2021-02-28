package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/appOLD/models/jorm/jdb"
	"github.com/oknors/okno/appOLD/models/jorm/x"
)

type BitTrexExchange struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketCurrency     string      `json:"MarketCurrency"`
		BaseCurrency       string      `json:"BaseCurrency"`
		MarketCurrencyLong string      `json:"MarketCurrencyLong"`
		BaseCurrencyLong   string      `json:"BaseCurrencyLong"`
		MinTradeSize       float64     `json:"MinTradeSize"`
		MarketName         string      `json:"MarketName"`
		IsActive           bool        `json:"IsActive"`
		IsRestricted       bool        `json:"IsRestricted"`
		Created            string      `json:"Created"`
		Notice             interface{} `json:"Notice"`
		IsSponsored        interface{} `json:"IsSponsored"`
		LogoURL            string      `json:"LogoUrl"`
	} `json:"result"`
}
type BitTrexExchangeTickers struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Result  []BitTrexExchangeTicker `json:"result"`
}
type BitTrexExchangeTicker struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	Volume         float64 `json:"Volume"`
	Last           float64 `json:"Last"`
	BaseVolume     float64 `json:"BaseVolume"`
	TimeStamp      string  `json:"TimeStamp"`
	Bid            float64 `json:"Bid"`
	Ask            float64 `json:"Ask"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
	PrevDay        float64 `json:"PrevDay"`
	Created        string  `json:"Created"`
}

func getBitTrexExchange() {
	fmt.Println("GetBitTrexExchangeStart")
	exchangeRaw := BitTrexExchange{}
	tickersRaw := BitTrexExchangeTickers{}

	slug := "bittrex"
	var exchange x.Exchange
	exchange.Name = "BitTrex"
	exchange.Slug = slug
	resps, err := http.Get("https://api.bittrex.com/api/v1.1/public/getmarkets")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://api.bittrex.com/api/v1.1/public/getmarketsummaries")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]BitTrexExchangeTicker)

	for _, ticker := range tickersRaw.Result {
		tickers[ticker.MarketName] = ticker
	}

	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw.Result {
		cur := x.Currency{
			Symbol: marketSrc.MarketCurrency,
			Ask:    fmt.Sprintf("%f", tickers[marketSrc.MarketName].Ask),
			Bid:    fmt.Sprintf("%f", tickers[marketSrc.MarketName].Bid),
			High:   fmt.Sprintf("%f", tickers[marketSrc.MarketName].High),
			Last:   fmt.Sprintf("%f", tickers[marketSrc.MarketName].Last),
			Low:    fmt.Sprintf("%f", tickers[marketSrc.MarketName].Low),
			Volume: fmt.Sprintf("%f", tickers[marketSrc.MarketName].Volume),
		}
		_, ok := markets[marketSrc.BaseCurrency]
		if !ok {
			markets[marketSrc.BaseCurrency] = []x.Currency{}
		}
		markets[marketSrc.BaseCurrency] = append(markets[marketSrc.BaseCurrency], cur)
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
	fmt.Println("GetBitTrexExchangeDone")

}
