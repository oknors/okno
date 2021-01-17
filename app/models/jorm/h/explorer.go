package h

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/jorm/a"
)

func ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	lastblock := a.RPCSRC(v["coin"]).GetBlockCount()
	lb := map[string]interface{}{
		"d": map[string]interface{}{
			"currentPage": page,
			"pageCount":   a.RPCSRC(v["coin"]).GetBlockCount() / per,
			"blocks":      a.RPCSRC(v["coin"]).GetBlocks(per, page),
			"lastBlock":   lastblock,
		},
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
	bl := map[string]interface{}{
		"d": lastblock,
	}
	out, err := json.Marshal(bl)
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

	bl := map[string]interface{}{
		"d": block,
	}
	fmt.Println("IP RPC source:", block)
	out, err := json.Marshal(bl)
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

	tX := a.RPCSRC(v["coin"]).GetTx(txid)

	tx := map[string]interface{}{
		"d": tX,
	}
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
	rmp := map[string]interface{}{
		"d": rawMemPool,
	}
	out, err := json.Marshal(rmp)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := a.RPCSRC(v["coin"]).GetMiningInfo()

	mi := map[string]interface{}{
		"d": miningInfo,
	}
	out, err := json.Marshal(mi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := a.RPCSRC(v["coin"]).GetInfo()

	in := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(in)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := a.RPCSRC(v["coin"]).GetPeerInfo()
	pi := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(pi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
