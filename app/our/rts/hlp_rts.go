package rts

import (
	"io/ioutil"
	"net/http"

	"github.com/oknors/okno/app/our/mod"
)

func getData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)
	return mapBody, err
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getCoin(coin string) mod.Coin {
	//gdb, err := jdb.OpenDB()
	//if err != nil {
	//}
	vCoin := mod.VCoin{}
	//if err := gdb.Read("coins", coin, &vCoin); err != nil {
	//	fmt.Println("Error", err)
	//}
	return vCoin.Coin
}
