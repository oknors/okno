package rts

import (

	"html/template"

	"net/http"


	"github.com/oknors/okno/app/our/tools"
)

var templates = make(map[string]*template.Template)

var last string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["404"] = template.Must(template.ParseFiles("app/our/tpl/hlp/plgs.gohtml", "app/our/tpl/css/boot.gohtml", "app/our/tpl/css/grid.gohtml", "app/our/tpl/css/typo.gohtml", "app/our/tpl/css/btn.gohtml", "app/our/tpl/hlp/base.gohtml", "app/our/tpl/hlp/body.gohtml", "app/our/tpl/hlp/head.gohtml", "app/our/tpl/hlp/style.gohtml", "app/our/tpl/hlp/spectre.gohtml", "app/our/tpl/404.gohtml", "app/our/tpl/hlp/search.gohtml"))

	// templates["cmc"] = template.Must(template.ParseFiles("tpl/frame/cmc.gohtml"))

	templates["ccw"] = template.Must(template.ParseFiles("app/our/tpl/frames/prices/ccw.gohtml"))
	templates["trades"] = template.Must(template.ParseFiles("app/our/tpl/frames/prices/trades.gohtml"))
	templates["nodes"] = template.Must(template.ParseFiles("app/our/tpl/frames/maps/nodes.gohtml"))
	templates["magnet"] = template.Must(template.ParseFiles("app/our/tpl/frames/games/magnet.gohtml"))
	templates["spiralogo"] = template.Must(template.ParseFiles("app/our/tpl/frames/anim/spiralogo.gohtml"))

	templates["proxy"] = template.Must(template.ParseFiles("app/our/tpl/proxy/proxy.gohtml"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := tools.GetData("index", "http://127.0.0.1:3553/")
	renderTemplate(w, "proxy", "proxy", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := "COINS"
	renderTemplate(w, "home", "base", data)
}

func FrameHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//frame := vars["frame"]
	//gdb, err := jdb.OpenDB()
	//if err != nil {
	//}
	//vCoin := mod.VCoin{}
	//if err := gdb.Read("coins", coin, &vCoin); err != nil {
	//	fmt.Println("Error", err)
	//}
	//nodesUrl := ComServer + "a/n/" + coin
	//gamp, err := http.Get(nodesUrl)
	//if err != nil {
	//	fmt.Println("AMP gampgampgampgamp", gamp)
	//}

	//fmt.Println("Read error", err)
	//defer gamp.Body.Close()
	//mapNodes, err := ioutil.ReadAll(gamp.Body)
	//gnodes := make(map[string]interface{})
	//json.Unmarshal(mapNodes, &gnodes)
	//if err != nil {
	//	fmt.Println("Read error", err)
	//}
	//
	//vCoin.Nodes = gnodes["nodes"]
	//data := vCoin
	//renderTemplate(w, frame, frame, data)
}
