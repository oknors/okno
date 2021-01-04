package app

import (
	"fmt"
	"time"
)

// Status
type Status struct {
	Working bool
	Coin    string
	Block   int
	Date    time.Time
}

type Node struct {
	ZIP         string
	WIP         string
	Name        string
	Domains     []string
	Description string
	LastSeen    time.Time
}

// Ticker
func (o *OKNO) Ticker() {
	//nodes :=  []Node{}
	//err := o.Database.Read("status", "nodes", &nodes)
	//if err != nil {
	//}
	//ticker := time.NewTicker(100 * time.Second)
	//for _ = range ticker.C {
	fmt.Println("tock")
	//for _, node := range nodes {
	//
	//
	//}
	//if err := o.Database.Create("status", "info", nodes); err != nil {
	//	fmt.Println("Error", err)
	//}
	//}
}
