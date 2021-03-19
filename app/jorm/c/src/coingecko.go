package csrc

import (
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/jorm/c"
	"github.com/oknors/okno/pkg/utl"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getCoinGecko() {
	fmt.Println("GetCoinGeckoStart")
	var coinsRaw []map[string]interface{}
	respcs, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	utl.ErrorLog(err)
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	if mapBody != nil {
		json.Unmarshal(mapBody, &coinsRaw)
		for _, coinSrc := range coinsRaw {
			if coinSrc["id"] != nil {
				if coinSrc["id"] != nil && coinSrc["id"].(string) != "" {
					slug := utl.MakeSlug(coinSrc["id"].(string))
					coin := c.Coin{}
					if coinSrc["name"] != nil {
						coin.Name = insertString(coinSrc["name"].(string), coin.Name)
					}
					if coinSrc["symbol"] != nil {
						coin.Ticker = insertString(coinSrc["symbol"].(string), coin.Ticker)
					}
					_, err = os.Stat(cfg.Path + "/jorm/coins/" + slug)
					if err != nil {
						getCoin(&coin, coinSrc["id"].(string), slug)
						fmt.Println("Insert Coin: ", coin.Name)

					} else {
						err = jdb.JDB.Read("jorm/coins", slug, &coin)
						utl.ErrorLog(err)
							if !coin.Checked["cg"] {
								fmt.Println("Ima Coin: ", coin.Name)
								if coin.Algo == "" || coin.StartDate == "" || coin.BlockTime == 0 || coin.Description == "" || coin.WebSite == nil || coin.Explorer  == nil || coin.Chat  == nil ||	coin.BitcoinTalk == "" ||	coin.Twitter == "" ||	coin.Facebook == "" ||	coin.Reddit == "" ||	coin.Logo.Img256 == "" {
									getCoin(&coin, coinSrc["id"].(string), slug)
									fmt.Println("Changed Coin: ", coin.Name)
								}
							}
						}
					}
			}
		}
	}
	fmt.Println("GetCoinGeckoDone")
}

func getCoin(coin *c.Coin, id, slug string) {
	fmt.Println("Checked1:", coin.Checked)
	if coin.Checked == nil {
			coin.Checked = make(map[string]bool)
		}
			coin.Slug = slug
			coinDetails := make(map[string]interface{})
			respc, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id + "?tickers=false&market_data=false&community_data=true&developer_data=false&sparkline=false")
			utl.ErrorLog(err)
			defer respc.Body.Close()
			mapBody, err := ioutil.ReadAll(respc.Body)
			if mapBody != nil {
				json.Unmarshal(mapBody, &coinDetails)
				if coinDetails["description"] != nil {
					coin.Description = insertString(coinDetails["description"].(map[string]interface{})["en"].(string), coin.Description)
				}
				if coinDetails["hashing_algorithm"] != nil {
					coin.Algo = insertString(coinDetails["hashing_algorithm"].(string), coin.Algo)
				}
				if coinDetails["genesis_date"] != nil {
					coin.StartDate = insertString(coinDetails["genesis_date"].(string), coin.StartDate)
				}
				if coinDetails["block_time_in_minutes"] != nil {
					coin.BlockTime = insertFloat(coinDetails["block_time_in_minutes"].(float64), coin.BlockTime)
				}
				checkItem(coinDetails["links"], "block_time_in_minutes", func(){coin.BlockTime = insertFloat(coinDetails["block_time_in_minutes"].(float64), coin.BlockTime)})
				checkItem(coinDetails["links"], "homepage", func(){coin.WebSite = insertStringSlice(stringSlice(coinDetails["links"].(map[string]interface{})["homepage"].([]interface{})), coin.WebSite)})
				checkItem(coinDetails["links"], "blockchain_site", func(){coin.Explorer = insertStringSlice(stringSlice(coinDetails["links"].(map[string]interface{})["blockchain_site"].([]interface{})), coin.Explorer)})
				checkItem(coinDetails["links"], "chat_url", func(){coin.Chat = insertStringSlice(stringSlice(coinDetails["links"].(map[string]interface{})["chat_url"].([]interface{})), coin.Chat)})
				checkItem(coinDetails["links"], "bitcointalk_thread_identifier", func(){coin.BitcoinTalk = insertString(fmt.Sprintf("%f", int(coinDetails["links"].(map[string]interface{})["bitcointalk_thread_identifier"].(float64))), coin.BitcoinTalk)})
				checkItem(coinDetails["links"], "twitter_screen_name", func(){coin.Twitter = insertString(coinDetails["links"].(map[string]interface{})["twitter_screen_name"].(string), coin.Twitter)})
				checkItem(coinDetails["links"], "telegram_channel_identifier", func(){coin.Telegram = insertString(coinDetails["links"].(map[string]interface{})["telegram_channel_identifier"].(string), coin.Telegram)})
				checkItem(coinDetails["links"], "subreddit_url", func(){coin.Reddit = insertString(coinDetails["links"].(map[string]interface{})["subreddit_url"].(string), coin.Reddit)})
				checkItem(coinDetails["image"], "large", func(){
						var cImgs utl.Images
						if coinDetails["image"].(map[string]interface{})["large"].(string) != "" && coinDetails["image"].(map[string]interface{})["large"].(string) != "missing_large.png" {
							cImgs = utl.GetIMG(coinDetails["image"].(map[string]interface{})["large"].(string), cfg.Path+"/static/coins/", coin.Slug)
						}
						if coin.Logo.Img256 == "" {
							coin.Logo = cImgs
						}
				})
				fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
				fmt.Println("Ubo:", coin.Name)
				fmt.Println("Checked1:", coin.Checked)
				coin.Checked["cg"] = true
				fmt.Println("Checked2:", coin.Checked)
				jdb.JDB.Write("jorm/coins", slug, coin)
		}
	time.Sleep(99 * time.Millisecond)
	return
}


func checkItem (item interface{}, cell string, set func()){
		if item != nil {
			if item.(map[string]interface{})[cell] != nil {
				set()
			}
		}
}
