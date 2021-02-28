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

type BitZExchange struct {
	Status    int                             `json:"status"`
	Msg       string                          `json:"msg"`
	Data      map[string]BitZExchangeCurrency `json:"data"`
	Time      int                             `json:"time"`
	Microtime string                          `json:"microtime"`
	Source    string                          `json:"source"`
}

type BitZExchangeCurrency struct {
	Symbol          string `json:"symbol"`
	QuoteVolume     string `json:"quoteVolume"`
	Volume          string `json:"volume"`
	PriceChange     string `json:"priceChange"`
	PriceChange24H  string `json:"priceChange24h"`
	AskPrice        string `json:"askPrice"`
	AskQty          string `json:"askQty"`
	BidPrice        string `json:"bidPrice"`
	BidQty          string `json:"bidQty"`
	Open            string `json:"open"`
	High            string `json:"high"`
	Low             string `json:"low"`
	Now             string `json:"now"`
	FirstID         int    `json:"firstId"`
	LastID          int    `json:"lastId"`
	DealCount       int    `json:"dealCount"`
	OrderBy         int    `json:"orderBy"`
	NumberPrecision int    `json:"numberPrecision"`
	PricePrecision  int    `json:"pricePrecision"`
	TradeAreaID     int    `json:"tradeAreaId"`
	Cny             string `json:"cny"`
	Usd             string `json:"usd"`
	Krw             string `json:"krw"`
	Jpy             string `json:"jpy"`
}

func getBitZExchange() {
	fmt.Println("GetBitZExchangeStart")
	marketsRaw := BitZExchange{}
	slug := "bitz"
	var exchange x.Exchange
	exchange.Name = "BitZ"
	exchange.Slug = slug
	respcs, err := http.Get("https://apiv2.bitz.com/Market/tickerall")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	markets := make(map[string][]x.Currency)
	for key, marketSrc := range marketsRaw.Data {
		m := strings.Split(key, "_")
		cur := x.Currency{
			Symbol: m[0],
			Ask:    marketSrc.AskPrice,
			Bid:    marketSrc.BidPrice,
			High:   marketSrc.High,
			Last:   marketSrc.Now,
			Low:    marketSrc.Low,
			Volume: marketSrc.Volume,
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
	fmt.Println("GetBitZExchangeDone")

}
