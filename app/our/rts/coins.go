package rts

import (


	"net/http"


	"github.com/oknors/okno/app/our/amp"
	"github.com/oknors/okno/app/our/mod"
)

func CoinHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//gdb, err := jdb.OpenDB()
	//if err != nil {
	//}
	//vCoin := mod.VCoin{}
	//if err := gdb.Read("coins", coin, &vCoin); err != nil {
	//	fmt.Println("Error", err)
	//}
	//gCoin := vCoin.Coin
	data := mod.CoinVw{
	//	Coin: gCoin,
	//	AMP:  amp.AmP(),
	}
	renderTemplate(w, "coin", "base", data)
}
func CoinsHandler(w http.ResponseWriter, r *http.Request) {
	data := mod.HCL{
		AMP: amp.AmP(),
	}
	renderTemplate(w, "coins", "base", data)
}

func CoinsMadnessHandler(w http.ResponseWriter, r *http.Request) {
	//coins := ser.GetCoins()
	data := mod.HCL{
		//Coins: coins,
		AMP:   amp.AmP(),
	}
	renderTemplate(w, "cmdns", "base", data)
}

func ImgHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//size := vars["size"]
	//gdb, err := jdb.OpenDB()
	//if err != nil {
	//}
	//vCoin := mod.VCoin{}
	//if err := gdb.Read("coins", coin, &vCoin); err != nil {
	//	fmt.Println("Error", err)
	//}
	//var img string
	//switch size {
	//case "16":
	//	img = vCoin.Coin.Imgs.Img16
	//case "32":
	//	img = vCoin.Coin.Imgs.Img32
	//case "64":
	//	img = vCoin.Coin.Imgs.Img64
	//case "128":
	//	img = vCoin.Coin.Imgs.Img128
	//case "256":
	//	img = vCoin.Coin.Imgs.Img256
	//}
	//encoded, _ := base64.StdEncoding.DecodeString(img)
	//w.Write(encoded)
}

func IcoHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]
	//gdb, err := jdb.OpenDB()
	//if err != nil {
	//}
	//vCoin := mod.VCoin{}
	//if err := gdb.Read("coins", coin, &vCoin); err != nil {
	//	fmt.Println("Error", err)
	//}
	//img := vCoin.Coin.Imgs.Img16
	//encoded, _ := base64.StdEncoding.DecodeString(img)
	//w.Write(encoded)
}

func FOFHandler(w http.ResponseWriter, r *http.Request) {
	data := mod.HCL{
		AMP: amp.AmP(),
	}
	renderTemplate(w, "404", "base", data)
}

func CertHandler(w http.ResponseWriter, r *http.Request) {
	//ser.GetNames()

}

func CoinsAMP(w http.ResponseWriter, r *http.Request) {
	//url := ComServer + "a/c/a"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func CoinsAMPimg(w http.ResponseWriter, r *http.Request) {
	//url := ComServer + "a/c/i"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}

func CoinsBNAMP(w http.ResponseWriter, r *http.Request) {
	//url := ComServer + "a/c/b"
	//data, _ := getData(url)
	//w.Write([]byte(data))
}
