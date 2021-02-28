package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/appOLD/models/jorm/h"
)

func (o *OKNO) jorm(r *mux.Router) {
	////////////////
	// jorm
	////////////////
	s := r.Host("jorm.okno.rs").Subrouter()
	s.StrictSlash(true)

	s.HandleFunc("/", h.HomeHandler)

	f := s.PathPrefix("/f").Subrouter()
	f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	a := s.PathPrefix("/a").Subrouter()
	a.HandleFunc("/coins", h.CoinsHandler).Methods("GET")
	a.HandleFunc("/{coin}/nodes", h.CoinNodesHandler).Methods("GET")
	a.HandleFunc("/{coin}/{nodeip}", h.NodeHandler).Methods("GET")

	b := s.PathPrefix("/b").Subrouter()
	b.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{blockheight}", h.ViewBlockHeight).Methods("GET")
	b.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	b.HandleFunc("/{coin}/mempool", h.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", h.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", h.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", h.ViewPeers).Methods("GET")
	b.HandleFunc("/{coin}/market", h.ViewMarket).Methods("GET")

	j := s.PathPrefix("/j").Subrouter()

	j.PathPrefix("/").Handler(h.ViewJSON())

	j.Headers("Access-Control-Allow-Origin", "*")

	e := s.PathPrefix("/e").Subrouter()
	//e.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	//e.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	e.HandleFunc("/{sec}/{coin}/{type}/{file}", h.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)
	//e.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	//e.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	e.Headers("Access-Control-Allow-Origin", "*")

}
