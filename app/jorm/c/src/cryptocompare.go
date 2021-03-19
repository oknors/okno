package csrc

import (
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/jorm/c"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/pkg/utl"
)

func getCryptoCompare() {
	fmt.Println("GetCryptoCompareStart")
	respcs, err := http.Get("https://min-api.cryptocompare.com/data/all/coinlist")
	utl.ErrorLog(err)
	defer respcs.Body.Close()
	if respcs != nil {
		coinsRaw := make(map[string]interface{})
		mapBody, err := ioutil.ReadAll(respcs.Body)
		json.Unmarshal(mapBody, &coinsRaw)
		if coinsRaw["Data"] != nil {
			for _, coinSrc := range coinsRaw["Data"].(map[string]interface{}) {
				if coinSrc != nil {
					cs := coinSrc.(map[string]interface{})
					if cs["CoinName"] != nil {
						slug := utl.MakeSlug(cs["Name"].(string))
						var coin c.Coin
						_, err = os.Stat(cfg.Path + "/jorm/coins/" + slug)
						if err != nil {
							fmt.Println("No " + cs["CoinName"].(string) + " in the system")
							var cImgs utl.Images
							if cs["ImageUrl"] != nil {
								imgurl := fmt.Sprint(cs["ImageUrl"].(string))
								if imgurl != "<nil>" {
									cImgs = utl.GetIMG("https://cryptocompare.com"+imgurl, cfg.Path+"/static/coins/", coin.Slug)
								}
							}
							coin.Name = cs["Name"].(string)
							coin.Ticker = cs["Symbol"].(string)
							coin.Slug = slug
							coin.Description = cs["Description"].(string)
							coin.Token = cs["AssetTokenStatus"].(string)
							coin.Algo = cs["Algorithm"].(string)
							coin.Proof = cs["ProofType"].(string)
							coin.Logo = cImgs
							fmt.Println("Insert Coin: ", coin.Name)
							jdb.JDB.Write("jorm/coins", slug, coin)
						}
					}
				}
			}
		}
	}
	fmt.Println("GetCryptoCompareDone")
}
