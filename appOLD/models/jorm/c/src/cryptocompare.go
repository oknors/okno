package csrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/appOLD/models/jorm/c"
	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/appOLD/models/jorm/jdb"
	"github.com/oknors/okno/pkg/utl"
)

type CryptoCompareCoins struct {
	Response string
	Message  string
	Data     map[string]struct {
		Id                   string
		Url                  string
		ImageUrl             string
		ContentCreatedOn     string
		Name                 string
		Symbol               string
		CoinName             string
		FullName             string
		Algorithm            string
		ProofType            string
		FullyPremined        string
		TotalCoinSupply      string
		BuiltOn              string
		SmartContractAddress string
		PreMinedValue        string
		TotalCoinsFreeFloat  string
		SortOrder            string
		Sponsored            string
		IsTrading            string
		TotalCoinsMined      string
		BlockNumber          string
		NetHashesPerSecond   string
		BlockReward          string
		BlockTime            string
	}
}
type CryptoCompareCoin struct {
	Response string `json:"Response"`
	Message  string `json:"Message"`
	Data     struct {
		SEO struct {
			PageTitle       string `json:"PageTitle"`
			PageDescription string `json:"PageDescription"`
			BaseURL         string `json:"BaseUrl"`
			BaseImageURL    string `json:"BaseImageUrl"`
			OgImageURL      string `json:"OgImageUrl"`
			OgImageWidth    string `json:"OgImageWidth"`
			OgImageHeight   string `json:"OgImageHeight"`
		} `json:"SEO"`
		General struct {
			ID                   string `json:"Id"`
			DocumentType         string `json:"DocumentType"`
			H1Text               string `json:"H1Text"`
			DangerTop            string `json:"DangerTop"`
			WarningTop           string `json:"WarningTop"`
			InfoTop              string `json:"InfoTop"`
			Symbol               string `json:"Symbol"`
			URL                  string `json:"Url"`
			BaseAngularURL       string `json:"BaseAngularUrl"`
			Name                 string `json:"Name"`
			ImageURL             string `json:"ImageUrl"`
			Description          string `json:"Description"`
			Features             string `json:"Features"`
			Technology           string `json:"Technology"`
			TotalCoinSupply      string `json:"TotalCoinSupply"`
			DifficultyAdjustment string `json:"DifficultyAdjustment"`
			BlockRewardReduction string `json:"BlockRewardReduction"`
			Algorithm            string `json:"Algorithm"`
			ProofType            string `json:"ProofType"`
			StartDate            string `json:"StartDate"`
			Twitter              string `json:"Twitter"`
			WebsiteURL           string `json:"WebsiteUrl"`
			Website              string `json:"Website"`
			Sponsor              struct {
				TextTop           string `json:"TextTop"`
				Link              string `json:"Link"`
				ImageURL          string `json:"ImageUrl"`
				ExcludedCountries string `json:"ExcludedCountries"`
			} `json:"Sponsor"`
			IndividualSponsor struct {
				Text              string `json:"Text"`
				ExcludedCountries string `json:"ExcludedCountries"`
				AffiliateLogo     string `json:"AffiliateLogo"`
				Link              string `json:"Link"`
				Type              string `json:"Type"`
			} `json:"IndividualSponsor"`
			LastBlockExplorerUpdateTS int     `json:"LastBlockExplorerUpdateTS"`
			BlockNumber               int     `json:"BlockNumber"`
			BlockTime                 int     `json:"BlockTime"`
			NetHashesPerSecond        float64 `json:"NetHashesPerSecond"`
			TotalCoinsMined           float64 `json:"TotalCoinsMined"`
			PreviousTotalCoinsMined   float64 `json:"PreviousTotalCoinsMined"`
			BlockReward               float64 `json:"BlockReward"`
		} `json:"General"`
		ICO struct {
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
		} `json:"ICO"`
		Subs            []string `json:"Subs"`
		StreamerDataRaw []string `json:"StreamerDataRaw"`
	} `json:"Data"`
	Type int `json:"Type"`
}

func getCryptoCompare() {
	fmt.Println("GetCryptoCompareStart")
	coinsRaw := CryptoCompareCoins{}
	respcs, err := http.Get("https://min-api.cryptocompare.com/data/all/coinlist")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &coinsRaw)
	for _, coinSrc := range coinsRaw.Data {
		slug := utl.MakeSlug(coinSrc.CoinName)
		var coin c.Coin
		if jdb.DB.Read(cfg.Web+"/coins", slug, &coin) != nil {
			var cData c.CoinData
			imgurl := fmt.Sprint(coinSrc.ImageUrl)
			ccID := coinSrc.Id
			coin.Name = coinSrc.CoinName
			coin.Ticker = coinSrc.Symbol
			coin.Slug = slug
			cData.Name = coinSrc.CoinName
			cData.Ticker = coinSrc.Symbol
			cData.Slug = slug
			if coinSrc.Algorithm != "" {
				cData.Algo = coinSrc.Algorithm
			}
			coinDetailsRaw := CryptoCompareCoin{}
			respc, err := http.Get("https://www.cryptocompare.com/api/data/coinsnapshotfullbyid/?id=" + ccID)
			if err != nil {
			}
			defer respc.Body.Close()
			mapBody, err := ioutil.ReadAll(respc.Body)
			json.Unmarshal(mapBody, &coinDetailsRaw)
			coinDetails := coinDetailsRaw.Data
			cData.Description = insertData(coinDetails.General.Description, cData.Description)
			cData.WebSite = insertData(coinDetails.General.WebsiteURL, cData.WebSite)
			cData.TotalCoinSupply = insertData(coinDetails.General.TotalCoinSupply, cData.TotalCoinSupply)
			cData.DifficultyAdjustment = insertData(coinDetails.General.DifficultyAdjustment, cData.DifficultyAdjustment)
			cData.DifficultyAdjustment = insertData(coinDetails.General.ProofType, cData.ProofType)
			cData.StartDate = insertData(coinDetails.General.StartDate, cData.StartDate)
			cData.Twitter = insertData(coinDetails.General.Twitter, cData.Twitter)
			fmt.Println("CryptoCompare Coin >>>", coin.Name)
			fmt.Println("CryptoCompare SlugSlugSlugSlug >>>", coin.Slug)
			fmt.Println("CryptoCompare ssssssssssssssssss >>>", slug)
			// fmt.Println("ImgUrl >>>", imgurl)
			// fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

			if coinDetails.ICO.Status != "N/A" {
				cData.Ico = true
				jdb.WriteCoinData(slug, "ico", coinDetails.ICO)

				fmt.Println("Insert ICO Coin: ", coinDetails.General.Name)

			}
			var cImgs utl.Images
			if imgurl != "<nil>" {
				cImgs, _ = utl.GetIMG("https://cryptocompare.com" + imgurl)
				if err != nil {
					fmt.Println("Problem Insert", err)
				}
			}
			jdb.WriteCoinImg(slug, cImgs)
			jdb.WriteCoin(slug, coin, cData)
		}
	}
	fmt.Println("GetCryptoCompareDone")
}
