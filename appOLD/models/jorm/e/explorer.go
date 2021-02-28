package e

import (
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/appOLD/models/jorm/a"
	"github.com/oknors/okno/appOLD/models/jorm/c"
	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/appOLD/models/jorm/jdb"
	"github.com/oknors/okno/appOLD/models/jorm/n"
	"github.com/oknors/okno/pkg/utl"
)

type Explorer struct {
	Status map[string]uint64 `json:"status"`
}

// GetExplorer updates the data from blockchain of a coin in the database
func GetExplorer(coins c.Coins) {
	var b []string
	for _, coin := range coins.C {
		var bn n.BitNoded
		www := cfg.Web + "data/" + coin.Slug
		if utl.FileExists(cfg.Dir + www + "/info/bitnodes") {
			b = append(b, coin.Slug)
			bitNodes := a.BitNodes{}
			if err := jdb.DB.Read(www+"/info", "bitnodes", &bitNodes); err != nil {
				fmt.Println("Error", err)
			}
			for _, bitnode := range bitNodes {
				err := GetBlockchain(cfg.Dir, www, bitnode)
				if err != nil {

				}
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
		}
	}

	//et := mod.Cache{Data: e}
	//jdb.DB.Write(cfg.Web, "explorer", et)
}

// GetExplorer returns the full set of information about a block
func GetBlockchain(dir, www string, a a.BitNode) (err error) {
	if utl.FileExists(dir + www + "/info/explorer") {
		e := Explorer{}
		if err := jdb.DB.Read(www+"/info", "explorer", &e); err != nil {
			fmt.Println("Error", err)
		}
		blockCount := a.GetBlockCount()
		if e.Status != nil && blockCount >= int(e.Status["blocks"]) {
			for {
				//e.block(a, www)
				fmt.Println("BlocksPre", e.Status["blocks"])
				blocksIndex := map[uint64]string{}
				if err := jdb.DB.Read(www+"/index", "blocks", &blocksIndex); err != nil {
					fmt.Println("Error", err)
				}
				e.Status["blocks"]++
				blockRaw := a.GetBlockByHeight(int(e.Status["blocks"]))
				if blockRaw != nil && blockRaw != "" {
					blockHash := blockRaw.(map[string]interface{})["hash"].(string)
					blocksIndex[e.Status["blocks"]] = blockHash
					go jdb.DB.Write(www+"/blocks", blockHash, blockRaw)
					block := (blockRaw).(map[string]interface{})
					if e.Status["blocks"] != 0 {
						for _, t := range (block["tx"]).([]interface{}) {
							e.tx(a, www, t.(string))
						}
					}
					//fmt.Println("BlocksPosle", e.Status["blocks"])
					jdb.DB.Write(www+"/info", "explorer", e)
					jdb.DB.Write(www+"/index", "blocks", blocksIndex)
				} else {
					break
				}
			}
		}
	} else {
		e := &Explorer{
			Status: map[string]uint64{"blocks": 0, "txs": 0, "addresses": 0},
		}
		jdb.DB.Write(www+"/info", "explorer", e)
		jdb.DB.Write(www+"/explorer/index", "blocks", map[uint64]string{})
		jdb.DB.Write(www+"/explorer/index", "txs", map[uint64]string{})
		jdb.DB.Write(www+"/explorer/index", "addresses", map[uint64]string{})
	}
	return
}

func (e *Explorer) tx(a a.BitNode, www, txid string) {
	txRaw := a.GetTx(txid)
	txsIndex := map[uint64]string{}
	if err := jdb.DB.Read(www+"/index", "txs", &txsIndex); err != nil {
		fmt.Println("Error", err)
	}
	e.Status["txs"]++
	txsIndex[e.Status["txs"]] = txid
	//fmt.Println("txid", txid)
	go jdb.DB.Write(www+"/txs", txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							e.addr(www, address.(string))
						}
					}
				}
			}
		}
	}
	jdb.DB.Write(www+"/index", "txs", txsIndex)
	return
}

func (e *Explorer) addr(www, address string) {
	addressesIndex := map[uint64]string{}
	if err := jdb.DB.Read(www+"/index", "addresses", &addressesIndex); err != nil {
		fmt.Println("Error", err)
	}
	//addressData := new(interface{})
	//if err := jdb.DB.Read(www, "explorer", &e); err != nil {
	//	fmt.Println("Error", err)
	//}
	addressData := address
	e.Status["addresses"]++
	addressesIndex[e.Status["addresses"]] = address
	go jdb.DB.Write(www+"/addresses", address, addressData)
	jdb.DB.Write(www+"/index", "addresses", addressesIndex)
	//fmt.Println("address", address)
	return
}

func (e *Explorer) status() {
	//s =
}
