package csrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/utl"
	"github.com/oknors/okno/app/models/jorm/c"
	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
)

type CoinCodexCoins []struct {
	Symbol                 string      `json:"symbol"`
	DisplaySymbol          string      `json:"display_symbol"`
	Name                   string      `json:"name"`
	Shortname              string      `json:"shortname"`
	LastPriceUsd           float64     `json:"last_price_usd"`
	PriceChange1HPercent   float64     `json:"price_change_1H_percent"`
	PriceChange1DPercent   float64     `json:"price_change_1D_percent"`
	PriceChange7DPercent   float64     `json:"price_change_7D_percent"`
	PriceChange30DPercent  float64     `json:"price_change_30D_percent"`
	PriceChange90DPercent  float64     `json:"price_change_90D_percent"`
	PriceChange180DPercent float64     `json:"price_change_180D_percent"`
	PriceChange365DPercent float64     `json:"price_change_365D_percent"`
	PriceChangeYTDPercent  float64     `json:"price_change_YTD_percent"`
	Volume24Usd            float64     `json:"volume_24_usd"`
	Display                string      `json:"display"`
	TradingSince           string      `json:"trading_since"`
	Supply                 int         `json:"supply"`
	Flags                  string      `json:"flags"`
	LastUpdate             string      `json:"last_update"`
	IcoEnd                 interface{} `json:"ico_end"`
	IncludeSupply          string      `json:"include_supply"`
	MarketCapUsd           float64     `json:"market_cap_usd"`
}

type CoinCodexCoin struct {
	Symbol                 string      `json:"symbol"`
	CoinName               string      `json:"coin_name"`
	Shortname              string      `json:"shortname"`
	Slug                   string      `json:"slug"`
	DisplaySymbol          string      `json:"display_symbol"`
	Display                string      `json:"display"`
	ReleaseDate            string      `json:"release_date"`
	IcoPrice               interface{} `json:"ico_price"`
	TodayOpen              float64     `json:"today_open"`
	Description            string      `json:"description"`
	PriceHigh24Usd         float64     `json:"price_high_24_usd"`
	PriceLow24Usd          float64     `json:"price_low_24_usd"`
	Start                  interface{} `json:"start"`
	End                    interface{} `json:"end"`
	IsPromoted             interface{} `json:"is_promoted"`
	Message                string      `json:"message"`
	Website                string      `json:"website"`
	Whitepaper             string      `json:"whitepaper"`
	TotalSupply            string      `json:"total_supply"`
	Supply                 int         `json:"supply"`
	Platform               string      `json:"platform"`
	HowToBuyURL            interface{} `json:"how_to_buy_url"`
	LastPriceUsd           float64     `json:"last_price_usd"`
	PriceChange1HPercent   string      `json:"price_change_1H_percent"`
	PriceChange1DPercent   string      `json:"price_change_1D_percent"`
	PriceChange7DPercent   string      `json:"price_change_7D_percent"`
	PriceChange30DPercent  string      `json:"price_change_30D_percent"`
	PriceChange90DPercent  string      `json:"price_change_90D_percent"`
	PriceChange180DPercent string      `json:"price_change_180D_percent"`
	PriceChange365DPercent string      `json:"price_change_365D_percent"`
	PriceChangeYTDPercent  string      `json:"price_change_YTD_percent"`
	Volume24Usd            float64     `json:"volume_24_usd"`
	TradingSince           string      `json:"trading_since"`
	StagesStart            interface{} `json:"stages_start"`
	StagesEnd              interface{} `json:"stages_end"`
	IncludeSupply          string      `json:"include_supply"`
	AthUsd                 string      `json:"ath_usd"`
	AthDate                string      `json:"ath_date"`
	NotTradingSince        interface{} `json:"not_trading_since"`
	LastUpdate             string      `json:"last_update"`
	Social                 interface{} `json:"social"`
	Socials                interface{} `json:"socials"`
}

func getCoinCodex() {
	fmt.Println("GetCoinCodexStart")
	coinsRaw := CoinCodexCoins{}
	respcs, err := http.Get("https://coincodex.com/apps/coincodex/cache/all_coins.json")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &coinsRaw)
	for _, coinSrc := range coinsRaw {

		slug := utl.MakeSlug(coinSrc.Name)
		var coin c.Coin
		if jdb.DB.Read(cfg.Web+"/coins", slug, &coin) != nil {
			var cData c.CoinData

			coin.Name = coinSrc.Name
			coin.Ticker = coinSrc.Symbol
			coin.Slug = slug
			cData.Name = coinSrc.Name
			cData.Ticker = coinSrc.Symbol
			coinDetails := CoinCodexCoin{}
			respc, err := http.Get("https://coincodex.com/api/coincodex/get_coin/" + coin.Ticker)
			if err != nil {
			}

			defer respc.Body.Close()
			mapBody, err := ioutil.ReadAll(respc.Body)
			json.Unmarshal(mapBody, &coinDetails)
			cData.Description = insertData(coinDetails.Description, cData.Description)
			cData.WebSite = insertData(coinDetails.Website, cData.WebSite)
			cData.TotalCoinSupply = insertData(coinDetails.TotalSupply, cData.TotalCoinSupply)
			cData.WhitePaper = insertData(coinDetails.Whitepaper, cData.WhitePaper)
			fmt.Println("CoinCodex Coin >>>", coin.Name)
			fmt.Println("CoinCodex SlugSlugSlugSlug >>>", coin.Slug)
			fmt.Println("CoinCodex ssssssssssssssssss >>>", slug)

			// fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

			if coinDetails.IcoPrice != "" {
				cData.Ico = true
				// jdb.WriteCoinData(slug, "ico", coinDetails.ICO)

				fmt.Println("Insert ICO Coin: ", coinDetails.CoinName)

			}
			var cImgs utl.Images

			cImgs, _ = utl.GetIMG("https://coincodex.com/en/resources/images/admin/coins/" + slug + ".png")
			if err != nil {
				fmt.Println("Problem Insert", err)
			}

			// fmt.Print("Desc >>>", coinDetails.Description)
			// fmt.Print("Coin >>>", coinSrc.Name)
			// // fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
			jdb.WriteCoinImg(slug, cImgs)
			jdb.WriteCoin(slug, coin, cData)
		}
	}

	fmt.Println("GetCoinCodexDone")

}
