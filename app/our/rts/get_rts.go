package rts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/our/mod"
)

func CoinNewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var newslist []mod.News
	nwd := make(map[string]map[string]map[string]interface{})

	url := "https://query.yahooapis.com/v1/public/yql?q=select%20title%2Clink%2C%20pubDate%20from%20rss%20where%20url%20%3D%20'https%3A%2F%2Fnews.google.com%2Fnews%2Frss%2Fsearch%2Fsection%2Fq%2F" + coin + "%3Fned%3Dus%26gl%3DUS%26hl%3Den'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
	data, _ := getData(url)
	json.Unmarshal(data, &nwd)

	nwq := nwd["query"]["results"]["item"]
	fmt.Println(reflect.TypeOf(nwq))

	switch mapNews := nwq.(type) {
	case []interface{}:
		for d := range mapNews {
			var nw = mapNews[d].(map[string]interface{})
			var news mod.News
			news = mod.News{
				Title: (nw["title"]).(string),
				Link:  (nw["link"]).(string),
				Date:  (nw["pubDate"]).(string),
			}
			newslist = append(newslist, news)
		}
	}
	fmt.Println("urlurlurlurlr", url)
	vnews := mod.VNews{
		News: newslist,
	}
	mnews, err := json.Marshal(vnews)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(mnews))
}
