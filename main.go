package main

import (
	"fmt"
	"github.com/oknors/okno/app"
	"log"
	"time"
)

func main() {
	//coins := c.Coins{}
	//coins := c.ReadAllCoins()

	okno := app.NewOKNO()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff

				fmt.Println("Radi kron GetBitNodes")
				//go e.GetExplorer(coins)
				//n.GetBitNodes(coins)
			//csrc.GetCoinSources()
			//xsrc.GetExchangeSources()
			// dsrc.GetDataSources()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	//exp.SrcNode().GetAddrs()

	//
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("Listening on port:")

	log.Fatal(okno.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
