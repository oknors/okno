package h

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/jorm/a"
	"net/http"
	"strconv"
)

func ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	lastblock := a.RPCSRC(v["coin"]).GetBlockCount()
	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   a.RPCSRC(v["coin"]).GetBlockCount() / per,
		"blocks":      a.RPCSRC(v["coin"]).GetBlocks(per, page),
		"lastBlock":   lastblock,
	}

	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func LastBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	lastblock := a.RPCSRC(v["coin"]).GetBlockCount()
	out, err := json.Marshal(lastblock)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var block interface{}
	id, err := strconv.ParseUint(v["id"], 10, 64)
	if err != nil {
		block = (a.RPCSRC(v["coin"]).GetBlock(v["id"])).(map[string]interface{})
	} else {
		block = a.RPCSRC(v["coin"]).GetBlockByHeight(int(id))
	}
	//fmt.Println("IP RPC source:", block)
	out, err := json.Marshal(block)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockheight"]
	// node := Node{}
	bhi, _ := strconv.Atoi(bh)
	block := a.RPCSRC(v["coin"]).GetBlockByHeight(bhi)
	out, err := json.Marshal(block)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewHash(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockhash"]
	block := (a.RPCSRC(v["coin"]).GetBlock(bh)).(map[string]interface{})
	h := strconv.FormatInt(block["height"].(int64), 10)
	http.Redirect(w, r, "/b/"+v["coin"]+"/block/"+h, 301)
}

func ViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	txid := v["txid"]
	// node := Node{}

	tx := a.RPCSRC(v["coin"]).GetTx(txid)

	out, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := a.RPCSRC(v["coin"]).GetRawMemPool()
	out, err := json.Marshal(rawMemPool)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := a.RPCSRC(v["coin"]).GetMiningInfo()

	out, err := json.Marshal(miningInfo)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := a.RPCSRC(v["coin"]).GetInfo()
	out, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := a.RPCSRC(v["coin"]).GetPeerInfo()
	out, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
