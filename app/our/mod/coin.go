package mod

import "github.com/oknors/okno/app/our/amp"

type Coin struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	Algo   string `json:"algo"`
	//Algo   Algo   `json:"algo"`
	BitNode bool     `json:"bitnode"`
	Imgs     Imgs     `json:"imgs"`
	CData    CoinData `json:"cdata"`
}
type Imgs struct {
	Img16  string `json:"img16"`
	Img32  string `json:"img32"`
	Img64  string `json:"img64"`
	Img128 string `json:"img128"`
	Img256 string `json:"img256"`
}

type Coins struct {
	Coins interface{} `json:"coins"`
}

type CoinData struct {
	Name                 string `json:"name"`
	Description          string `json:"desc"`
	WebsiteUrl           string `json:"web"`
	TotalCoinSupply      string `json:"total"`
	DifficultyAdjustment string `json:"diff"`
	BlockRewardReduction string `json:"rew"`
	ProofType            string `json:"proof"`
	StartDate            string `json:"start"`
	Twitter              string `json:"tw"`
}

// type Algo struct {
// 	A string `json:"a"`
// 	Z string `json:"z"`
// }

type HCL struct {
	Coins []VCoin `json:"coins"`
	//Algos []Algo        `json:"algos"`
	AMP amp.AMP `json:"amp"`
}
type VCoin struct {
	Coin  Coin        `json:"coin"`
	Nodes interface{} `json:"nodes"`
}

type CoinVw struct {
	Coin Coin `json:"coin"`
	//Algos []Algo        `json:"algos"`
	AMP amp.AMP `json:"amp"`
}

type CoinAmp struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	Algo   string `json:"algo"`
	BitNode bool     `json:"bitnode"`
}
type BCoinAmp struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	Algo   string `json:"algo"`
}
