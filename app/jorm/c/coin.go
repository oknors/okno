package c

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/pkg/utl"

	"image"
	"strings"

	"github.com/oknors/okno/app/jdb"
)

type Coins struct {
	N int    `json:"n"`
	C []Coin `json:"c"`
}
type CoinsBase struct {
	N int        `json:"n"`
	C []CoinBase `json:"c"`
}

// Coin stores identifying information about coins in the database
type CoinBase struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
	Img    string `json:"i"`
}

type Coin struct {
	Name                 string          `json:"name" form:"name"`
	Ticker               string          `json:"ticker" form:"ticker"`
	Slug                 string          `json:"slug" form:"slug"`
	Algo                 string          `json:"algo" form:"algo"`
	Selected             bool            `json:"selected" form:"selected"`
	Favorite             bool            `json:"fav" form:"favorite"`
	Token                string          `json:"token" form:"token"`
	Platform             string          `json:"platform" form:"platform"`
	Proof                string          `json:"proof" form:"proof"`
	Description          string          `json:"description" form:"description"`
	Rank                 int             `json:"rank" form:"rank"`
	CoinName             string          `json:"coinname" form:"coinname"`
	BitNode              bool            `json:"bitnode" form:"bitnode"`
	Ico                  bool            `json:"ico" form:"ico"`
	PlatformType         string          `json:"type"`
	TotalCoinSupply      float64         `json:"total"`
	BuiltOn              string          `json:"builton"`
	BlockTime            float64         `json:"blocktime"`
	DifficultyAdjustment string          `json:"diff"`
	BlockRewardReduction string          `json:"rew"`
	ProofType            string          `json:"proof"`
	StartDate            string          `json:"start"`
	WebSite              []string        `json:"web"`
	Explorer             []string        `json:"explorer"`
	Chat                 []string        `json:"chat"`
	Twitter              string          `json:"tw"`
	Facebook             string          `json:"facebook"`
	Telegram             string          `json:"telegram"`
	Reddit               string          `json:"reddit"`
	Github               []string        `json:"github"`
	BitcoinTalk          string          `json:"bitcointalk"`
	WhitePaper           string          `json:"whitepaper"`
	Logo                 utl.Images      `json:"logo" form:"logo"`
	Published            bool            `json:"published" form:"published"`
	Checked              map[string]bool `json:"checked"`
}



// ReadAllCoins reads in all of the data about all coins in the database
func ReadAllCoins() Coins {
	csb := LoadCoinsBase(false)
	cns := Coins{
		N: csb.N,
		C: getCoins(),
	}
	jdb.JDB.Write("jorm/info", "coinsbase", csb)

	jdb.JDB.Write("jorm/info", "coins", LoadCoinsBase(true))
	return cns
}

func (coin *Coin) SelectCoin() *Coin {
	//coin.LogoBig = LoadLogo(coin.Slug, "img128")
	//coin.Data = LoadInfo(coin.Slug)
	return coin
}
func LoadLogo(slug, size string) image.Image {
	// Load logo image from database
	logos := make(map[string]interface{})
	fmt.Println("slug", slug)
	err := jdb.JDB.Read("jorm/data/"+slug, "logo", logos)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
	logo, _, err := image.Decode(reader)
	utl.ErrorLog(err)
	return logo
}

func LoadInfo(slug string) Coin {
	// Load coin data from database
	info := Coin{}
	err := jdb.JDB.Read("data/"+slug, "info", info)
	utl.ErrorLog(err)
	//jsonString, _ := json.Marshal(info)

	// convert json to struct
	//s := CoinData{}
	//json.Unmarshal(jsonString, &s)
	return info
}

func LoadCoinsBase(filter bool) CoinsBase {
	coins := getCoins()
	csb := CoinsBase{}
	csb.N = 0
	for i, coin := range coins {
		ccb := CoinBase{
			Rank:   csb.N,
			Name:   coin.Name,
			Ticker: coin.Ticker,
			Slug:   coin.Slug,
			Algo:   coin.Algo,
		}
		if filter {
			if coins[i].Platform != "token" && coins[i].Algo != "" && coins[i].Algo != "N/A" && coins[i].Proof != "N/A" {
				ccb.Img = coin.Logo.Img16
				csb.N++
				csb.C = append(csb.C, ccb)
			}
		} else {
			csb.N++
			csb.C = append(csb.C, ccb)
		}

	}
	return csb
}

func getCoins() []Coin {
	data, err := jdb.JDB.ReadAll("jorm/coins")
	utl.ErrorLog(err)
	coins := make([][]byte, len(data))
	for i := range data {
		coins[i] = []byte(data[i])
	}
	cs := make([]Coin, len(coins))
	for i := range coins {
		if err := json.Unmarshal(coins[i], &cs[i]); err != nil {
			fmt.Println("Error", err)
		}
	}
	return cs
}
