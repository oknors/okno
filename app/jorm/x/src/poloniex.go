package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/jorm/x"
)

type PoloniexExchange map[string]struct {
	ID            int    `json:"id"`
	Last          string `json:"last"`
	LowestAsk     string `json:"lowestAsk"`
	HighestBid    string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
	IsFrozen      string `json:"isFrozen"`
	High24Hr      string `json:"high24hr"`
	Low24Hr       string `json:"low24hr"`
}

func getPoloniexExchange() {
	fmt.Println("GetPoloniexExchangeStart")
	marketsRaw := PoloniexExchange{}
	slug := "poloniex"
	var exchange x.Exchange
	exchange.Name = "Poloniex"
	exchange.Slug = slug
	respcs, err := http.Get("https://poloniex.com/public?command=returnTicker")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for key, marketSrc := range marketsRaw {
		m := strings.Split(key, "_")
		cur := x.Currency{
			Symbol: m[1],
			Ask:    marketSrc.LowestAsk,
			Bid:    marketSrc.HighestBid,
			High:   marketSrc.High24Hr,
			Last:   marketSrc.Last,
			Low:    marketSrc.Low24Hr,
			Volume: marketSrc.BaseVolume,
		}
		_, ok := markets[m[0]]
		if !ok {
			markets[m[0]] = []x.Currency{}
		}
		markets[m[0]] = append(markets[m[0]], cur)
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

	fmt.Println("GetPoloniexExchangeDone")

}
