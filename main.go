package main

import (
	"fmt"
	"github.com/oknors/okno/app"
	"github.com/oknors/okno/app/cfg"
	//"github.com/oknors/okno/app/jorm/coin"
	//"github.com/oknors/okno/app/jorm/exchange"
	"log"
	//"time"
)

func main() {
	okno := app.NewOKNO()
	//coins := coin.ReadAllCoins()
	//_ = exchange.ReadAllExchanges()
	//ticker := time.NewTicker(999 * time.Second)
	//quit := make(chan struct{})
	//go func() {
	//	for {
	//		select {
			//case <-ticker.C:
				//app.Tickers(coins)
				//fmt.Println("OKNO wooikos")
			//case <-quit:
			//	ticker.Stop()
			//	return
			//}
		//}
	//}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("Listening on port: ", cfg.CONFIG.Port)
	log.Fatal(okno.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
