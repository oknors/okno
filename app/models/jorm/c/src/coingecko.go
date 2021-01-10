package csrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/oknors/okno/app/utl"
	"github.com/oknors/okno/app/models/jorm/c"
	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
)

type CoinGeckoCoins []struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoinGeckoCoin struct {
	ID                 string   `json:"id"`
	Symbol             string   `json:"symbol"`
	Name               string   `json:"name"`
	BlockTimeInMinutes int      `json:"block_time_in_minutes"`
	Categories         []string `json:"categories"`
	Description        struct {
		En string `json:"en"`
	} `json:"description"`
	Links struct {
		Homepage                    []string    `json:"homepage"`
		BlockchainSite              []string    `json:"blockchain_site"`
		OfficialForumURL            []string    `json:"official_forum_url"`
		ChatURL                     []string    `json:"chat_url"`
		AnnouncementURL             []string    `json:"announcement_url"`
		TwitterScreenName           string      `json:"twitter_screen_name"`
		FacebookUsername            string      `json:"facebook_username"`
		BitcointalkThreadIdentifier int         `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string      `json:"telegram_channel_identifier"`
		SubredditURL                interface{} `json:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `json:"github"`
			Bitbucket []interface{} `json:"bitbucket"`
		} `json:"repos_url"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	CountryOrigin       string      `json:"country_origin"`
	GenesisDate         interface{} `json:"genesis_date"`
	MarketCapRank       int         `json:"market_cap_rank"`
	CoingeckoRank       int         `json:"coingecko_rank"`
	CoingeckoScore      float64     `json:"coingecko_score"`
	DeveloperScore      float64     `json:"developer_score"`
	CommunityScore      float64     `json:"community_score"`
	LiquidityScore      float64     `json:"liquidity_score"`
	PublicInterestScore float64     `json:"public_interest_score"`
	MarketData          struct {
		TotalSupply       float64   `json:"total_supply"`
		CirculatingSupply float64   `json:"circulating_supply"`
		LastUpdated       time.Time `json:"last_updated"`
	} `json:"market_data"`
	CommunityData struct {
		FacebookLikes            int         `json:"facebook_likes"`
		TwitterFollowers         int         `json:"twitter_followers"`
		RedditAveragePosts48H    float64     `json:"reddit_average_posts_48h"`
		RedditAverageComments48H float64     `json:"reddit_average_comments_48h"`
		RedditSubscribers        int         `json:"reddit_subscribers"`
		RedditAccountsActive48H  int         `json:"reddit_accounts_active_48h"`
		TelegramChannelUserCount interface{} `json:"telegram_channel_user_count"`
	} `json:"community_data"`
	DeveloperData struct {
		Forks                        int `json:"forks"`
		Stars                        int `json:"stars"`
		Subscribers                  int `json:"subscribers"`
		TotalIssues                  int `json:"total_issues"`
		ClosedIssues                 int `json:"closed_issues"`
		PullRequestsMerged           int `json:"pull_requests_merged"`
		PullRequestContributors      int `json:"pull_request_contributors"`
		CodeAdditionsDeletions4Weeks struct {
			Additions interface{} `json:"additions"`
			Deletions interface{} `json:"deletions"`
		} `json:"code_additions_deletions_4_weeks"`
		CommitCount4Weeks              int           `json:"commit_count_4_weeks"`
		Last4WeeksCommitActivitySeries []interface{} `json:"last_4_weeks_commit_activity_series"`
	} `json:"developer_data"`
	PublicInterestStats struct {
		AlexaRank   int `json:"alexa_rank"`
		BingMatches int `json:"bing_matches"`
	} `json:"public_interest_stats"`
	StatusUpdates []interface{} `json:"status_updates"`
	LastUpdated   time.Time     `json:"last_updated"`
	Tickers       []struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Market struct {
			Name                string `json:"name"`
			Identifier          string `json:"identifier"`
			HasTradingIncentive bool   `json:"has_trading_incentive"`
		} `json:"market"`
		Last          float64 `json:"last"`
		Volume        float64 `json:"volume"`
		ConvertedLast struct {
			Btc float64 `json:"btc"`
			Eth float64 `json:"eth"`
			Usd float64 `json:"usd"`
		} `json:"converted_last"`
		ConvertedVolume struct {
			Btc float64 `json:"btc"`
			Eth float64 `json:"eth"`
			Usd float64 `json:"usd"`
		} `json:"converted_volume"`
		TrustScore             string    `json:"trust_score"`
		BidAskSpreadPercentage float64   `json:"bid_ask_spread_percentage"`
		Timestamp              time.Time `json:"timestamp"`
		LastTradedAt           time.Time `json:"last_traded_at"`
		LastFetchAt            time.Time `json:"last_fetch_at"`
		IsAnomaly              bool      `json:"is_anomaly"`
		IsStale                bool      `json:"is_stale"`
		TradeURL               string    `json:"trade_url"`
		CoinID                 string    `json:"coin_id"`
	} `json:"tickers"`
}

func getCoinGecko() {
	fmt.Println("GetCoinGeckoStart")
	coinsRaw := CoinGeckoCoins{}
	respcs, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
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
			coinDetails := CoinGeckoCoin{}
			respc, err := http.Get("https://api.coingecko.com/api/v3/coins/" + coinSrc.ID)
			if err != nil {
			}
			defer respc.Body.Close()
			mapBody, err := ioutil.ReadAll(respc.Body)
			json.Unmarshal(mapBody, &coinDetails)
			cData.Description = insertData(coinDetails.Description.En, cData.Description)
			cData.WebSite = insertData(coinDetails.Links.Homepage[0], cData.WebSite)
			cData.Twitter = insertData(coinDetails.Links.TwitterScreenName, cData.Twitter)
			cData.Facebook = insertData(coinDetails.Links.FacebookUsername, cData.Facebook)
			// insertData(coinDetails.Links.ReposURL.Github[0], cData.Facebook)
			fmt.Println("CoinGecko Coin >>>", coin.Name)
			fmt.Println("CoinGecko SlugSlugSlugSlug >>>", coin.Slug)
			fmt.Println("CoinGecko ssssssssssssssssss >>>", slug)

			// fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

			// if coinDetails.IcoPrice != "" {
			// 	cData.Ico = true
			// 	// jdb.WriteCoinData(slug, "ico", coinDetails.ICO)

			// 	fmt.Println("Insert ICO Coin: ", coinDetails.CoinName)

			// }
			var cImgs utl.Images
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", coinDetails.Image.Large)
			if coinDetails.Image.Large != "" && coinDetails.Image.Large != "missing_large.png" {
				cImgs, _ = utl.GetIMG(coinDetails.Image.Large)
				if err != nil {
					fmt.Println("Problem Insert", err)
				}
			}
			// fmt.Print("Desc >>>", coinDetails.Description)
			// fmt.Print("Coin >>>", coinSrc.Name)
			// // fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
			jdb.WriteCoinImg(slug, cImgs)
			jdb.WriteCoin(slug, coin, cData)
		}
	}
	fmt.Println("GetCoinGeckoDone")

}
