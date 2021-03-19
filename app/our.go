package app

import (
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/pkg/utl"
	"net/http"

	//"github.com/oknors/okno/app/jorm/h"
	//"github.com/oknors/okno/app/our/rts"


)

func (o *OKNO) our(mr *mux.Router) {
	////////////////
	// our
	////////////////
	s := mr.Host("com-http.us").Subrouter()
	//s.StrictSlash(true)
	//s.HandleFunc("/", h.HomeHandler)
	host :=Host{
		Name: "COM-HTTP",
		Slug: "comhttpus",
		Template: "comhttp",
		Host: "com-http.us",
	}

	var hashes []string

	hashes = append(hashes,  "sha384-" + utl.SHA384(cfg.Path+"/templates/comhttp/tpl/js/vue.js"))

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		title := host.Name
		data := map[string]interface{}{
			"Title":    title,
			"Site":     "com-http",
			"Section":  "index",
			"Hashes": hashes,
		}
		o.template(w, host, data)
	})
	//s.Path("/").HandlerFunc(rts.NXIndexHandler).Name("nxindex")
	//s.Path("/coins").HandlerFunc(rts.NXCoinsHandler).Name("nxcoins")
	//s.Path("/words").HandlerFunc(rts.NXWordsHandler).Name("nxwords")
	//s.Path("/a/coins").HandlerFunc(rts.CoinsAMP).Name("coinsamp")
	//s.Path("/a/coinsimg").HandlerFunc(rts.CoinsAMPimg).Name("coinsampimg")
	//s.Path("/a/bitnodes").HandlerFunc(rts.CoinsBNAMP).Name("coinsampbitnodes")
	//
	c := mr.Host("{coin}.com-http.us").Subrouter()
	c.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		coin := mux.Vars(r)["coin"]
		title := host.Name
		section := "coin"
		data := map[string]interface{}{
			"Title":    title,
			"Slug":     coin,
			"Section":  section,
			"Hashes": hashes,
		}
		o.template(w, host, data)
	})
	//c.Path("/").HandlerFunc(rts.NXCoinHandler).Name("nxcoin")
	//
	//
	//c.Path("/explorer").HandlerFunc(rts.NXExplorerHandler).Name("nxexplorer")
	//c.Path("/explorer/block/{id}").HandlerFunc(rts.NXBlockHandler).Name("nxblock")
	//c.Path("/explorer/hash/{id}").HandlerFunc(rts.NXHashHandler).Name("nxhash")
	//c.Path("/explorer/tx/{id}").HandlerFunc(rts.NXTxHandler).Name("nxtx")
	//c.Path("/explorer/addr/{id}").HandlerFunc(rts.NXAddrsHandler).Name("nxaddr")
	//
	//c.Path("/explorer/search").HandlerFunc(rts.DoSearch).Name("search")
	//
	//c.Path("/network").HandlerFunc(rts.NXNetworkHandler).Name("nxnetwork")
	//c.Path("/price").HandlerFunc(rts.NXPriceHandler).Name("nxprice")
	//c.Path("/ecosystem").HandlerFunc(rts.NXEcoHandler).Name("nxeco")
	//s.Host("f.com-http.us").Path("/{frame}/{file}").HandlerFunc(rts.Frames).Name("nxframes")
	//
	//c.Path("/a/b").HandlerFunc(rts.ApiLastBlock).Name("b")
	//c.Path("/a/bta/{id}").HandlerFunc(rts.ApiBlockTxAddr).Name("bta")
	//
	//c.Path("/a/i").HandlerFunc(rts.ApiInfo).Name("info")
	//c.Path("/a/p").HandlerFunc(rts.ApiPeer).Name("peer")
	//c.Path("/a/m").HandlerFunc(rts.ApiMiningInfo).Name("mining")
	//c.Path("/a/r").HandlerFunc(rts.ApiRawMemPool).Name("rawmempool")
	//
	//c.Path("/a/n").HandlerFunc(rts.NodesHandler).Name("nodes")
	//
	//c.Path("/a/{type}/{id}").HandlerFunc(rts.ApiData).Name("coin")
	//
	//c.Path("/a/n").HandlerFunc(rts.CoinNewsHandler).Name("news")
	//c.Path("/f/cmc").HandlerFunc(rts.CMCHandler).Name("cmc")
	//
	//c.Path("/favicon.ico").HandlerFunc(rts.IcoHandler).Name("ico")
	//
	//s.Host("i.com-http.us").Path("/{coin}/{size}").HandlerFunc(rts.ImgHandler).Name("img")
	//
	//c.Path("/frames/{frame}").HandlerFunc(rts.FrameHandler).Name("frame")
	//
	//s.PathPrefix("/json/").Handler(http.StripPrefix("/json/", http.FileServer(http.Dir("./JDB/"))))
	//
	//s.Host("l.com-http.us").PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/libs/"))))
	//
	//s.NotFoundHandler = http.HandlerFunc(rts.FOFHandler)
	//s.Schemes("https")


	//a := s.PathPrefix("/a").Subrouter()
	s.Headers("Access-Control-Allow-Origin", "*")
	c.Headers("Access-Control-Allow-Origin", "*")

}

