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

type LBankExchange []struct {
	Symbol string `json:"symbol"`
	Ticker struct {
		Change   int     `json:"change"`
		High     float64 `json:"high"`
		Latest   float64 `json:"latest"`
		Low      float64 `json:"low"`
		Turnover float64 `json:"turnover"`
		Vol      float64 `json:"vol"`
	} `json:"ticker"`
	Timestamp int64 `json:"timestamp"`
}

func getLBankExchange() {
	fmt.Println("GetLBankExchangeStart")
	marketsRaw := LBankExchange{}
	slug := "lbank"
	var exchange x.Exchange
	exchange.Name = "LBank"
	exchange.Slug = slug
	respcs, err := http.Get("https://api.lbkex.com/v1/ticker.do?symbol=all")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range marketsRaw {
		m := strings.Split(marketSrc.Symbol, "_")
		cur := x.Currency{
			Symbol: strings.ToUpper(m[0]),
			// Ask:    marketSrc.AskPrice,
			// Bid:    marketSrc.BidPrice,
			High:   fmt.Sprintf("%f", marketSrc.Ticker.High),
			Last:   fmt.Sprintf("%f", marketSrc.Ticker.Latest),
			Low:    fmt.Sprintf("%f", marketSrc.Ticker.Low),
			Volume: fmt.Sprintf("%f", marketSrc.Ticker.Vol),
		}
		_, ok := markets[strings.ToUpper(m[1])]
		if !ok {
			markets[strings.ToUpper(m[1])] = []x.Currency{}
		}
		markets[strings.ToUpper(m[1])] = append(markets[strings.ToUpper(m[1])], cur)
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

	fmt.Println("GetLBankExchangeDone")

}
