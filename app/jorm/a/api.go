package a

import (
	"fmt"

	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/pkg/utl"
)

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNodeSrc is a node's address
type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"p"`
	Jrc  *utl.Endpoint
}

func RPCSRC(c string) (b *BitNode) {
	bitNodes := BitNodes{}
	if err := jdb.JDB.Read("jorm/conf/"+c, "bitnodes", &bitNodes); err != nil {
		fmt.Println("Errdor", err)
	}
	for _, bn := range bitNodes {
		b = &bn
		b.Jrc = utl.NewClient(cfg.CONFIG.RPC.Username, cfg.CONFIG.RPC.Password, b.IP, b.Port)
		if b.Jrc != nil {
			fmt.Println("bitb:", b.IP)
			break
		}

	}
	fmt.Println("b:", b)
	return
}
