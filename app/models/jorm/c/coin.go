package c

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"image"
	"strings"

	"github.com/oknors/okno/app/models/jorm"
	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
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
	Rank int    `json:"r"`
	Name string `json:"n"`
	Slug string `json:"s"`
}

type Coin struct {
	Name     string `json:"n" form:"name"`
	Ticker   string `json:"t" form:"ticker"`
	Slug     string `json:"s" form:"slug"`
	Selected bool
	Favorite bool
	Logo     image.Image `json:"l" form:"logo"`
	LogoBig  image.Image `json:"lb" form:"logobig"`
	//Link     *p9.Clickable
	Data CoinData
}

// CoinData stores all of the information relating to a coin
type CoinData struct {
	Rank    int    `json:"rank" form:"rank"`
	Name    string `json:"name" form:"name"`
	Ticker  string `json:"symbol" form:"symbol"`
	Slug    string `json:"slug" form:"slug"`
	Algo    string `json:"algo" form:"algo"`
	BitNode bool   `json:"bitnode" form:"bitnode"`

	Token bool `json:"token" form:"token"`
	Ico   bool `json:"ico" form:"ico"`

	Description          string `json:"desc"`
	WebSite              string `json:"web"`
	TotalCoinSupply      string `json:"total"`
	DifficultyAdjustment string `json:"diff"`
	BlockRewardReduction string `json:"rew"`
	ProofType            string `json:"proof"`
	StartDate            string `json:"start"`

	Twitter string `json:"tw"`
	// Explorers            []string `json:"explorers"`
	// Boards               []string `json:"boards"`
	Facebook   string `json:"facebook"`
	Reddit     string `json:"reddit"`
	Github     string `json:"github"`
	WhitePaper string `json:"whitepaper"`
	Published  bool   `json:"published" form:"published"`
}

type ICO struct {
	Status                      string `json:"Status"`
	Description                 string `json:"Description"`
	TokenType                   string `json:"TokenType"`
	Website                     string `json:"Website"`
	PublicPortfolioURL          string `json:"PublicPortfolioUrl"`
	PublicPortfolioID           string `json:"PublicPortfolioId"`
	Features                    string `json:"Features"`
	FundingTarget               string `json:"FundingTarget"`
	FundingCap                  string `json:"FundingCap"`
	ICOTokenSupply              string `json:"ICOTokenSupply"`
	TokenSupplyPostICO          string `json:"TokenSupplyPostICO"`
	TokenPercentageForInvestors string `json:"TokenPercentageForInvestors"`
	TokenReserveSplit           string `json:"TokenReserveSplit"`
	Date                        int    `json:"Date"`
	EndDate                     int    `json:"EndDate"`
	FundsRaisedList             string `json:"FundsRaisedList"`
	FundsRaisedUSD              string `json:"FundsRaisedUSD"`
	StartPrice                  string `json:"StartPrice"`
	StartPriceCurrency          string `json:"StartPriceCurrency"`
	PaymentMethod               string `json:"PaymentMethod"`
	Jurisdiction                string `json:"Jurisdiction"`
	LegalAdvisers               string `json:"LegalAdvisers"`
	LegalForm                   string `json:"LegalForm"`
	SecurityAuditCompany        string `json:"SecurityAuditCompany"`
	Blog                        string `json:"Blog"`
	WhitePaper                  string `json:"WhitePaper"`
	WhitePaperLink              string `json:"WhitePaperLink"`
}

// ReadAllCoins reads in all of the data about all coins in the database
func ReadAllCoins() Coins {
	coins := jdb.ReadData(cfg.Web + "/coins")
	cs := make([]Coin, len(coins))
	csb := CoinsBase{}

	csb.N = 0
	for i := range coins {
		csb.N++
		if err := json.Unmarshal(coins[i], &cs[i]); err != nil {
			fmt.Println("Error", err)
		}

		ccb := CoinBase{
			Rank: csb.N,
			Name: cs[i].Name,
			Slug: cs[i].Slug,
		}
		cs[i].Logo = LoadLogo(cs[i].Slug, "img32")
		//cs[i].Link = new(p9.Clickable)
		fmt.Println("ccb", ccb)

		csb.C = append(csb.C, ccb)
	}

	cns := Coins{
		N: csb.N,
		C: cs,
	}
	c := mod.Cache{Data: cns}
	cb := mod.Cache{Data: csb}
	jdb.DB.Write(cfg.Web, "coins", c)
	jdb.DB.Write(cfg.Web, "coinsbase", cb)
	return cns
}

func (coin *Coin) SelectCoin() *Coin {
	coin.LogoBig = LoadLogo(coin.Slug, "img128")
	coin.Data = LoadInfo(coin.Slug)
	return coin
}
func LoadLogo(slug, size string) image.Image {
	// Load logo image from database
	fmt.Println("slug", slug)

	l := jdb.Read("data/"+slug, "logo")
	logos := l.(map[string]interface{})
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
	logo, _, err := image.Decode(reader)
	if err != nil {
		//log.Fatal(err)
	}
	return logo
}

func LoadInfo(slug string) CoinData {
	// Load coin data from database
	info := jdb.Read("data/"+slug, "info")
	jsonString, _ := json.Marshal(info)
	//fmt.Println(string(jsonString))
	// convert json to struct
	s := CoinData{}
	json.Unmarshal(jsonString, &s)
	return s
}
