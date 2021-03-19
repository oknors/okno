package rts

import (

	"net/http"

	"github.com/oknors/okno/app/our/amp"
	"github.com/oknors/okno/app/our/mod"
	"github.com/gorilla/mux"
)

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func ExplorerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	gCoin := getCoin(coin)
	data := mod.Explorer{
		Coin: gCoin,
		AMP:  amp.AmP(),
	}
	renderTemplate(w, "explorer", "base", data)
}

func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.BlVw{
		ID:    id,
		Coin:  gCoin,
		Block: mod.Block{},
		AMP:   amp.AmP(),
	}
	//fmt.Println("datadatadatadatadata", data)

	renderTemplate(w, "blockheight", "base", data)
}
func ViewBlockHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.BlVw{
		ID:    id,
		Coin:  gCoin,
		Block: mod.Block{},
		AMP:   amp.AmP(),
	}
	//fmt.Println("datadatadatadatadata", data)

	renderTemplate(w, "blockhash", "base", data)
}
func ViewTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.TxVw{
		ID:   id,
		Coin: gCoin,
		Tx:   mod.Tx{},
		AMP:  amp.AmP(),
	}
	renderTemplate(w, "tx", "base", data)
}
func ViewAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	id := vars["id"]
	gCoin := getCoin(coin)
	data := mod.AdVw{
		ID:   id,
		Coin: gCoin,
		Addr: mod.Addr{},
		AMP:  amp.AmP(),
	}
	renderTemplate(w, "addr", "base", data)
}

func ApiData(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//id := vars["id"]
	//tp := vars["type"]
	//url := ComServer + "a/e/" + coin + "/" + tp + "/" + id
	//data, _ := getData(url)
	//w.Write([]byte(data))
}
func ApiLast(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/last"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func ApiLastBlock(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/b"
	//data, _ := getData(url)
	//w.Write([]byte(data))/
}

func ApiBlockTxAddr(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//id := vars["id"]
	//url := ComServer + "a/e/" + coin + "/b/" + id
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/info"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}
func ApiPeer(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/peer"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}
func ApiMiningInfo(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/gmi"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}
func ApiRawMemPool(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/e/" + coin + "/rmp"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func NodesHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//url := ComServer + "a/n/" + coin
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func DoSearch(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//r.ParseForm()
	//search := r.FormValue("chsrc")
	//fmt.Println("searchsearchsearchsearchsearchsearchsearchsearch", search)
	//
	//tps := []string{"block", "hash", "tx", "addr"}
	//var tpt = "noresults"
	//for _, tp := range tps {
	//	url := ComServer + "a/e/" + coin + "/" + tp + "/" + search
	//	fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", tp)
		//resp, err := getData(url)
		//var search map[string]interface{}
		//json.Unmarshal(resp, &search)
		//if err != nil {
		//	fmt.Println("Read error", err)
		//}
		//if search["d"] != nil {
		//	tpt = tp
		//
		//}

	//}

	//http.Redirect(w, r, fmt.Sprintf("/explorer/"+tpt+"/"+search), http.StatusPermanentRedirect)
}
