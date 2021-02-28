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

type IDAXExchange struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Ticker    []struct {
		Pair string `json:"pair"`
		Open string `json:"open"`
		High string `json:"high"`
		Low  string `json:"low"`
		Last string `json:"last"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}

func getIDAXExchange() {
	fmt.Println("GetIDAXExchangeStart")
	marketsRaw := IDAXExchange{}
	slug := "idax"
	var exchange x.Exchange
	exchange.Name = "IDAX"
	exchange.Slug = slug
	respcs, err := http.Get("https://openapi.idax.pro//api/v2/ticker")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for _, marketSrc := range marketsRaw.Ticker {
		m := strings.Split(marketSrc.Pair, "_")
		cur := x.Currency{
			Symbol: m[0],
			// Ask: marketSrc.Low,
			// Bid: marketSrc.High,
			High:   marketSrc.High,
			Last:   marketSrc.Last,
			Low:    marketSrc.Low,
			Volume: marketSrc.Vol,
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

	fmt.Println("GetIDAXExchangeDone")

}
