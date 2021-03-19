package a

import (
	"fmt"
	"strconv"
	"time"
)

func (rpc *BitNode) GetBlockCount() (b int) {
	bparams := []int{}
	gbc, err := rpc.Jrc.MakeRequest("getblockcount", bparams)
	if err != nil {
		fmt.Println("Error n call: ", err)

	}
	switch gbc.(type) {
	case float64:
		return int(gbc.(float64))
	case string:
		b, _ := strconv.Atoi(gbc.(string))
		return b
	default:
		b, _ := strconv.Atoi(gbc.(string))
		return b
	}
	return
}

func (rpc *BitNode) GetBlocks(per, page int) (blocks []map[string]interface{}) {
	blockCount := rpc.GetBlockCount()
	fmt.Println("blockCount", blockCount)

	startBlock := blockCount - per*page
	minusBlockStart := int(startBlock + per)

	for ibh := minusBlockStart; ibh >= startBlock; {
		blocks = append(blocks, rpc.GetBlockShortByHeight(ibh))
		ibh--
	}
	return blocks
}

func (rpc *BitNode) GetBlock(blockhash string) (b interface{}) {
	bparams := []string{blockhash}
	b, err := rpc.Jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	//b := rawb.(map[string]interface{})
	//block := make(map[string]interface{})
	//if b["bits"] != nil {
	//	block["bits"] = b["bits"].(string)
	//}
	//if b["chainwork"] != nil {
	//	block["chainwork"] = b["chainwork"].(string)
	//}
	//if b["confirmations"] != nil {
	//	block["confirmations"] = int64(b["confirmations"].(float64))
	//}
	//if b["difficulty"] != nil {
	//	block["difficulty"] = b["difficulty"].(float64)
	//}
	//if b["hash"] != nil {
	//	block["hash"] = b["hash"].(string)
	//}
	//if b["height"] != nil {
	//	block["height"] = int64(b["height"].(float64))
	//}
	//if b["mediantime"] != nil {
	//	block["mediantime"] = int64(b["mediantime"].(float64))
	//}
	//if b["merkleroot"] != nil {
	//	block["merkleroot"] = b["merkleroot"].(string)
	//}
	//if b["nTx"] != nil {
	//	block["ntx"] = int(b["nTx"].(float64))
	//}
	//if b["nextblockhash"] != nil {
	//	block["nextblockhash"] = b["nextblockhash"].(string)
	//}
	//if b["nonce"] != nil {
	//	block["nonce"] = int64(b["nonce"].(float64))
	//}
	//if b["previousblockhash"] != nil {
	//	block["previousblockhash"] = b["previousblockhash"].(string)
	//}
	//if b["size"] != nil {
	//	block["size"] = int64(b["size"].(float64))
	//}
	//if b["strippedsize"] != nil {
	//	block["strippedsize"] = int64(b["strippedsize"].(float64))
	//}
	//if b["time"] != nil {
	//	unixTimeUTC := time.Unix(int64(b["time"].(float64)), 0)
	//	block["time"] = unixTimeUTC.Format(time.RFC850)
	//	block["timeutc"] = unixTimeUTC.Format(time.RFC3339)
	//}
	//if b["tx"] != nil {
	//	txsRaw := b["tx"].([]interface{})
	//	var txs []string
	//	for _, t := range txsRaw {
	//		txs = append(txs, t.(string))
	//	}
	//	block["tx"] = txs
	//}
	//if b["version"] != nil {
	//	block["version"] = int64(b["version"].(float64))
	//}
	//if b["versionHex"] != nil {
	//	block["versionhex"] = b["versionHex"].(string)
	//}
	//if b["weight"] != nil {
	//	block["weight"] = int64(b["weight"].(float64))
	//}
	//
	//if b["pow_algo"] != nil {
	//	block["pow_algo"] = b["pow_algo"].(string)
	//}
	//if b["pow_hash"] != nil {
	//	block["pow_hash"] = b["pow_hash"].(string)
	//}
	//if b["pow_algo_id"] != nil {
	//	block["pow_algo_id"] = int64(b["pow_algo_id"].(float64))
	//}

	return
}

func (rpc *BitNode) GetBlockShort(blockhash string) map[string]interface{} {
	bparams := []string{blockhash}
	rawb, err := rpc.Jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	b := rawb.(map[string]interface{})
	block := make(map[string]interface{})
	if b["bits"] != nil {
		block["bits"] = b["bits"].(string)
	}
	if b["confirmations"] != nil {
		block["confirmations"] = int64(b["confirmations"].(float64))
	}
	if b["difficulty"] != nil {
		block["difficulty"] = b["difficulty"].(float64)
	}
	if b["hash"] != nil {
		block["hash"] = b["hash"].(string)
	}
	if b["height"] != nil {
		block["height"] = int64(b["height"].(float64))
	}
	if b["tx"] != nil {
		var txsNumber int
		for _ = range b["tx"].([]interface{}) {
			txsNumber++
		}
		block["txs"] = txsNumber
	}
	if b["size"] != nil {
		block["size"] = int64(b["size"].(float64))
	}
	if b["time"] != nil {
		unixTimeUTC := time.Unix(int64(b["time"].(float64)), 0)
		block["time"] = unixTimeUTC.Format(time.RFC850)
		block["timeutc"] = unixTimeUTC.Format(time.RFC3339)
	}
	return block
}
func (rpc *BitNode) GetBlockShortByHeight(blockheight int) (block map[string]interface{}) {
	bparams := []int{blockheight}
	blockHash, err := rpc.Jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	if blockHash != nil {
		block = rpc.GetBlockShort((blockHash).(string))
	}
	return block
}

func (rpc *BitNode) GetBlockByHeight(blockheight int) (block interface{}) {
	bparams := []int{blockheight}
	blockHash, err := rpc.Jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	if blockHash != nil {
		block = rpc.GetBlock((blockHash).(string))
	}
	return block
}
