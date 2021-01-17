package e

import (
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/app/models/jorm/a"
	"github.com/oknors/okno/app/models/jorm/c"
	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/models/jorm/n"
	"github.com/oknors/okno/app/utl"
	"strconv"
)

type Explorer struct {
	Blocks int    `json:"blocks"`
	URL    string `json:"url"`
}

// GetExplorer updates the data from blockchain of a coin in the database
func GetExplorer(coins c.Coins) {
	var b []string
	//var bns n.BitNodeds
	for _, coin := range coins.C {
		var bn n.BitNoded
		www := cfg.Web + "data/" + coin.Slug
		if utl.FileExists(cfg.Dir + www + "/bitnodes.json") {
			b = append(b, coin.Slug)
			bitNodes := a.BitNodes{}
			if err := jdb.DB.Read(www, "bitnodes", &bitNodes); err != nil {
				fmt.Println("Error", err)
			}
			for _, bitnode := range bitNodes {
				bitNode := GetBlockchain(cfg.Dir, www, bitnode)
				//var dataBitNode = mod.Cache{
				//	Response: true,
				//	Data:     bitNode,
				//}
				//jdb.DB.Write(cfg.Web+"/bitnodes/", bitnode.IP, dataBitNode)

				fmt.Println("--------------------")
				fmt.Println("explorer", bitNode)
				fmt.Println("--------------------")

				bn.Coin = coin
			}
			//bns = append(bns, bn)
			nodes := jdb.ReadData(cfg.Web + "/nodes")
			ns := make(n.Nodes, len(nodes))

			for i := range nodes {
				if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
					fmt.Println("Error", err)
				}
			}
			//var dataNodes = mod.Cache{
			//	Response: true,
			//	Data:     ns,
			//}
			//jdb.DB.Write(cfg.Web, "nodes", dataNodes)
		}
	}

	//et := mod.Cache{Data: e}
	//jdb.DB.Write(cfg.Web, "explorer", et)
}

// GetExplorer returns the full set of information about a block
func GetBlockchain(dir, www string, a a.BitNode) (err error) {
	//getInfo := a.GetInfo()
	if utl.FileExists(dir + www + "/explorer.json") {
		e := Explorer{}
		if err := jdb.DB.Read(www, "explorer", &e); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("eeeeeeeeeee", e)

		for {
			blockRaw := a.GetBlockByHeight(e.Blocks)
			if blockRaw != "" {
				jdb.DB.Write(www+"/explorer/blocks", strconv.Itoa(e.Blocks), blockRaw)
				block := (blockRaw).(map[string]interface{})

				for _, t := range (block["tx"]).([]interface{}) {
					txid := t.(string)
					txRaw := a.GetTx(txid)
					jdb.DB.Write(www+"/explorer/txs", txid, txRaw)
					tx := (txRaw).(map[string]interface{})
					if tx["vout"] != nil {
						for _, nRaw := range tx["vout"].([]interface{}) {
							if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
								scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
								if scriptPubKey["addresses"] != nil {
									for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
										e := new(interface{})
										if err := jdb.DB.Read(www, "explorer", &e); err != nil {
											fmt.Println("Error", err)
										}
										jdb.DB.Write(www+"/explorer/addresses", address.(string), address)
										fmt.Println("address", address)
									}
								}
							}
							//fmt.Println("n", nRaw.(map[string]interface{})["n"])
							//fmt.Println("value", nRaw.(map[string]interface{})["value"])
						}
					}
				}

				//fmt.Println("--------------------")
				//fmt.Println("block", e.Blocks)
				//fmt.Println("cccc", www)
				e.Blocks = e.Blocks + 1
				jdb.DB.Write(www, "explorer", e)
			} else {
				break
			}
		}
	} else {
		e := Explorer{
			Blocks: 0,
		}
		jdb.DB.Write(www, "explorer", e)
	}
	return
}
