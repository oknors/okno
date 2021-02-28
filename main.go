package main

import (
	"fmt"
	"github.com/oknors/okno/app"
	"github.com/oknors/okno/appOLD/models/jorm/c"
	"log"
	"time"
)

func main() {
	okno := app.NewOKNO()
	coins := c.ReadAllCoins()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				app.Tickers(coins)
				fmt.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("Listening on port: ", okno.Configuration.AppPort)
	log.Fatal(okno.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
