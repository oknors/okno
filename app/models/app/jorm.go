package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/jorm/h"
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

	// s := r.PathPrefix("/s").Subrouter()
	// s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./tpl/static/"))))

	b := s.PathPrefix("/b").Subrouter()
	b.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	// b.HandleFunc("/{coin}/block/{blockheight}", h.ViewBlockHeight).Methods("GET")
	b.HandleFunc("/{coin}/block/{blockheight}", h.ViewBlockHeight).Methods("GET")
	b.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	b.HandleFunc("/{coin}/market", h.ViewMarket).Methods("GET")

	j := s.PathPrefix("/j").Subrouter()
	// j.HandleFunc("/", h.ViewJSON).Methods("GET")
	//j.Headers("X-Requested-With", "XMLHttpRequest")
	//j.Headers("X-Requested-With", "XMLHttpRequest")
	//j.Headers("X-Requested-With", "XMLHttpRequest")

	//j.Headers("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//j.Headers("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

	j.PathPrefix("/").Handler(h.ViewJSON())

	j.Headers("Access-Control-Allow-Origin", "*")

}
