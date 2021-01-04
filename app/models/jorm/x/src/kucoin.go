package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/models/jorm/x"
)

type KuCoinExchange struct {
	Code string `json:"code"`
	Data struct {
		Ticker []KuCoinExchangeCurrency `json:"ticker"`
		Time   int64                    `json:"time"`
	} `json:"data"`
}
type KuCoinExchangeCurrency struct {
	Symbol       string `json:"symbol"`
	High         string `json:"high,omitempty"`
	Vol          string `json:"vol"`
	Last         string `json:"last"`
	Low          string `json:"low,omitempty"`
	Buy          string `json:"buy"`
	Sell         string `json:"sell"`
	ChangePrice  string `json:"changePrice,omitempty"`
	AveragePrice string `json:"averagePrice,omitempty"`
	ChangeRate   string `json:"changeRate"`
	VolValue     string `json:"volValue"`
}

func getKuCoinExchange() {
	fmt.Println("GetKuCoinExchangeStart")
	marketsRaw := KuCoinExchange{}
	slug := "kucoin"
	var exchange x.Exchange
	exchange.Name = "KuCoin"
	exchange.Slug = slug
	respcs, err := http.Get("https://api.kucoin.com/api/v1/market/allTickers")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range marketsRaw.Data.Ticker {
		m := strings.Split(marketSrc.Symbol, "-")
		cur := x.Currency{
			Symbol: m[0],
			Ask:    marketSrc.Sell,
			Bid:    marketSrc.Buy,
			High:   marketSrc.High,
			Last:   marketSrc.Last,
			Low:    marketSrc.Low,
			Volume: marketSrc.VolValue,
		}
		_, ok := markets[m[1]]
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

	fmt.Println("GetKuCoinExchangeDone")

}
