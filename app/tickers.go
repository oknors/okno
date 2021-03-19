package app

import (
	"fmt"
	"github.com/oknors/okno/app/jorm/c"
	"github.com/oknors/okno/app/jorm/n"
)

func Tickers(coins c.Coins) {
	//coins := c.Coins{}

	fmt.Println("Cron is wooikos")
	//go e.GetExplorer(coins)
	go n.GetBitNodes(coins)
	//go csrc.GetCoinSources()
	//xsrc.GetExchangeSources()
	// dsrc.GetDataSources()
}
