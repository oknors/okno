package rts

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/our/tools"
	)

// func NXIndexHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.URL.Path = "/"
// 		p.ServeHTTP(w, r)
// 	}
// }

func NXIndexHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:3553/"
	data := tools.GetData("index", url)
	renderTemplate(w, "proxy", "proxy", data)
}

func NXCoinsHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:3553/coins/"
	data := tools.GetData("coins", url)
	renderTemplate(w, "proxy", "proxy", data)
}

func NXWordsHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:3553/x/words/"
	data := tools.GetData("xWords", url)
	renderTemplate(w, "proxy", "proxy", data)
}

func NXCoinHandler(w http.ResponseWriter, r *http.Request) {
	name := "S" + mux.Vars(r)["coin"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["coin"]
	data := tools.GetData(name, url)
	renderTemplate(w, "proxy", "proxy", data)
}

