package a

import (
	"fmt"

	"github.com/oknors/okno/appOLD/models/jorm/cfg"
	"github.com/oknors/okno/appOLD/models/jorm/jdb"
	"github.com/oknors/okno/pkg/utl"
)

// BitNodes is array of bitnodes addresses
type BitNodes []BitNode

// BitNodeSrc is a node's address
type BitNode struct {
	IP   string `json:"ip"`
	Port int64  `json:"port"`
	Jrc  *utl.Endpoint
}

func RPCSRC(c string) (b *BitNode) {
	bitNodes := BitNodes{}
	if err := jdb.DB.Read(cfg.Web+"/data/"+c+"/info", "bitnodes", &bitNodes); err != nil {
		fmt.Println("Errdor", err)
	}
	for _, bn := range bitNodes {
		b = &bn
		b.Jrc = utl.NewClient(cfg.Credentials.Username, cfg.Credentials.Password, b.IP, b.Port)
		if b.Jrc != nil {
			fmt.Println("bitb:", b.IP)
		}

	}
	fmt.Println("b:", b)
	return
}
