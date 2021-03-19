package rts

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/our/tools"
	)



func NXExplorerHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "EX" + mux.Vars(r)["id"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/explorer/"
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}
func NXBlockHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "BL" + mux.Vars(r)["id"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/explorer/block/" + mux.Vars(r)["id"]
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}

func NXHashHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "HASH" + mux.Vars(r)["id"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/explorer/hash/" + mux.Vars(r)["id"]
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}
func NXTxHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "TX" + mux.Vars(r)["id"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/explorer/tx/" + mux.Vars(r)["id"]
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}
func NXAddrsHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "ADDR" + mux.Vars(r)["id"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/explorer/addr/" + mux.Vars(r)["id"]
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}

func NXNetworkHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "Network"
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/network/"
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}
func NXPriceHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "Price"
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/price/"
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}
func NXEcoHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["coin"] + "EcoSystem"
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"] + "/ecosystem/"
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}

