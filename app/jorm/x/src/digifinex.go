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

type DigiFinexExchange struct {
	Ticker []struct {
		Vol     float64 `json:"vol"`
		Change  float64 `json:"change"`
		BaseVol float64 `json:"base_vol"`
		Sell    float64 `json:"sell"`
		Last    float64 `json:"last"`
		Symbol  string  `json:"symbol"`
		Low     float64 `json:"low"`
		Buy     float64 `json:"buy"`
		High    float64 `json:"high"`
	} `json:"ticker"`
	Date int `json:"date"`
	Code int `json:"code"`
}

func getDigiFinexExchange() {
	fmt.Println("GetDigiFinexExchangeStart")
	marketsRaw := DigiFinexExchange{}
	slug := "digifinex"
	var exchange x.Exchange
	exchange.Name = "DigiFinex"
	exchange.Slug = slug
	respcs, err := http.Get("https://openapi.digifinex.vip/v3/ticker")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range marketsRaw.Ticker {
		m := strings.Split(marketSrc.Symbol, "_")
		cur := x.Currency{
			Symbol: strings.ToUpper(m[0]),
			Ask:    fmt.Sprintf("%f", marketSrc.Sell),
			Bid:    fmt.Sprintf("%f", marketSrc.Buy),
			High:   fmt.Sprintf("%f", marketSrc.High),
			Last:   fmt.Sprintf("%f", marketSrc.Last),
			Low:    fmt.Sprintf("%f", marketSrc.Low),
			Volume: fmt.Sprintf("%f", marketSrc.Vol),
		}
		_, ok := markets[m[1]]
		if !ok {
			markets[m[1]] = []x.Currency{}
		}
		markets[m[1]] = append(markets[m[1]], cur)
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

	fmt.Println("GetDigiFinexExchangeDone")

}
