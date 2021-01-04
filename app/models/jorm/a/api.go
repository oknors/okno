package a

import (
	"fmt"

	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
)

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNodeSrc is a node's address
type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"port"`
}

func RPCSRC(c string) (b *BitNode) {
	bitNodes := BitNodes{}
	if err := jdb.DB.Read(cfg.Web+"/data/"+c, "bitnodes", &bitNodes); err != nil {
		fmt.Println("Errdor", err)
	}
	for _, bn := range bitNodes {
		// Need function for best node
		// {
		b = &bn
		fmt.Println("bitNodeSRC:", b)
		// }
	}
	return
}
