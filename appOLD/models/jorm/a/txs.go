package a

import (
	"fmt"

	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/pkg/utl"
)

func (rpc *BitNode) GetTx(txid string) (t interface{}) {
	jrc := utl.NewClient(cfg.Credentials.Username, cfg.Credentials.Password, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	verbose := int(1)
	var grtx []interface{}
	grtx = append(grtx, txid)
	grtx = append(grtx, verbose)
	t, err := jrc.MakeRequest("getrawtransaction", grtx)
	if err != nil {
		fmt.Println("Jorm Node Get Tx Error", err)
	}

	//t := rawt.(map[string]interface{})
	//tx := make(map[string]interface{})
	//
	//if t["blockhash"] != nil {
	//	tx["blockhash"] = t["blockhash"].(string)
	//}
	//if t["chainwork"] != nil {
	//	tx["chainwork"] = t["chainwork"].(string)
	//}
	//if t["blocktime"] != nil {
	//	tx["blocktime"] = int64(t["blocktime"].(float64))
	//}
	//if t["confirmations"] != nil {
	//	tx["confirmations"] = int64(t["confirmations"].(float64))
	//}
	//if t["difficulty"] != nil {
	//	tx["difficulty"] = t["difficulty"].(float64)
	//}
	//if t["hash"] != nil {
	//	tx["hash"] = t["hash"].(string)
	//}
	//if t["hex"] != nil {
	//	tx["hex"] = t["hex"].(string)
	//}
	//if t["locktime"] != nil {
	//	tx["locktime"] = int64(t["locktime"].(float64))
	//}
	//if t["size"] != nil {
	//	tx["size"] = int64(t["size"].(float64))
	//}
	//if t["time"] != nil {
	//	tx["time"] = int64(t["time"].(float64))
	//}
	//if t["txid"] != nil {
	//	tx["txid"] = t["txid"].(string)
	//}
	//if t["version"] != nil {
	//	tx["version"] = int64(t["version"].(float64))
	//}
	//if t["vin"] != nil {
	//	tx["vin"] = t["vin"].([]interface{})
	//}
	//if t["vout"] != nil {
	//	tx["vout"] = t["vout"].([]interface{})
	//}
	//if t["vsize"] != nil {
	//	tx["vsize"] = int64(t["vsize"].(float64))
	//}
	//if t["weight"] != nil {
	//	tx["weight"] = int64(t["weight"].(float64))
	//}
	return
}

func (rpc *BitNode) GetBlockTxAddr(blockheight int) interface{} {
	jrc := utl.NewClient(cfg.Credentials.Username, cfg.Credentials.Password, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{blockheight}
	blockHash, err := jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Tx Addr Error", err)
	}
	var block interface{}
	var txs []interface{}
	if blockHash != nil {
		block = rpc.GetBlock((blockHash).(string))
	}

	iblock := make(map[string]interface{})
	iblock = block.(map[string]interface{})
	itxs := iblock["tx"].([]interface{})
	for _, itx := range itxs {
		var txid string
		txid = itx.(string)

		verbose := int(1)
		var grtx []interface{}
		grtx = append(grtx, txid)
		grtx = append(grtx, verbose)
		rtx, err := jrc.MakeRequest("getrawtransaction", grtx)
		if err != nil {
			fmt.Println("Jorm Node Get Block Tx Addr Tx Error", err)
		}
		txs = append(txs, rtx)

	}
	blocktxaddr := map[string]interface{}{
		"b": block,
		"t": txs,
	}
	return blocktxaddr
}
