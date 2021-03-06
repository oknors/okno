package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
	"strings"

	//"net/url"

	"github.com/sensepost/gowitness/chrome"
	"github.com/sensepost/gowitness/lib"
	"github.com/sensepost/gowitness/storage"
	//"strings"
)

var (
	options = lib.NewOptions()
	chrm    = chrome.NewChrome()
	db      = storage.NewDb()
)

func (o *OKNO) img(r *mux.Router) {
	////////////////
	// img
	////////////////
	a := r.Host("img.okno.rs").Subrouter()
	a.StrictSlash(true)

	a.HandleFunc("/", o.viewWebImg).Methods("GET")

	//a.Headers("Access-Control-Allow-Origin", "*")
}

func (o *OKNO) viewWebImg(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimSpace(r.URL.Query().Get("url"))
	path := o.Configuration.Path + "/static/img"
	_, err := os.Stat(path + "/" + url)
	if err != nil {
		fmt.Println(path + "/" + url)
		fmt.Println("nema")
		exec.Command("gowitness single --destination " + path + " -o " + url + ".png https://" + url)
	} else {
		fmt.Println(path)
		fmt.Println("ima")
	}
	http.ServeFile(w, r, path+"/"+url)
}
