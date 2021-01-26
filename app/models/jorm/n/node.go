package n

import (
	"encoding/json"
	"fmt"

	"github.com/oknors/okno/app/models/jorm/a"
	"github.com/oknors/okno/app/models/jorm/c"
	"github.com/oknors/okno/app/models/jorm/cfg"
	"github.com/oknors/okno/app/models/jorm/jdb"
	"github.com/oknors/okno/app/utl"
)

type BitNodeds []BitNoded

// BitNoded data
type BitNoded struct {
	Coin     c.Coin          `json:"coin"`
	BitNodes []BitNodeStatus `json:"bitnodes"`
}

// NodeStatus stores current data for a node
type BitNodeStatus struct {
	Live           bool        `json:"live"`
	GetInfo        interface{} `json:"getinfo"`
	GetPeerInfo    interface{} `json:"getpeerinfo"`
	GetRawMemPool  interface{} `json:"getrawmempool"`
	GetMiningInfo  interface{} `json:"getmininginfo"`
	GetNetworkInfo interface{} `json:"getnetworkinfo"`
	GeoIP          interface{} `json:"geoip"`
}

type Nodes []NodeInfo

// NodeInfo stores info retrieved via geoip about a node
type NodeInfo struct {
	IP            string `json:"ip"`
	Port          int64  `json:"port"`
	Host          string `json:"host"`
	Rdns          string `json:"rdns"`
	ASN           string `json:"asn"`
	ISP           string `json:"isp"`
	CountryName   string `json:"countryname"`
	CountryCode   string `json:"countrycode"`
	RegionName    string `json:"regionname"`
	RegionCode    string `json:"regioncode"`
	City          string `json:"city"`
	Postcode      string `json:"postcode"`
	ContinentName string `json:"continentname"`
	ContinentCode string `json:"continentcode"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Zipcode       string `json:"zipcode"`
	Timezone      string `json:"timezone"`
	LastSeen      string `json:"lastseen"`
	Live          bool   `json:"live"`
}

// GetBitNodes updates the data about all of the coins in the database
func GetBitNodes(coins c.Coins) {
	var b []string
	var bns BitNodeds
	for _, coin := range coins.C {
		var bn BitNoded
		www := cfg.Dir + cfg.Web + "data/" + coin.Slug + "/"
		if utl.FileExists(www + "/info/bitnodes") {
			b = append(b, coin.Slug)
			bitNodes := a.BitNodes{}
			if err := jdb.DB.Read(cfg.Web+"data/"+coin.Slug, "bitnodes", &bitNodes); err != nil {
				fmt.Println("Error", err)
			}
			for _, bitnode := range bitNodes {

				bitNode := GetBitNodeStatus(bitnode)
				nds := GetNodes(bitnode)
				for _, n := range nds {
					jdb.DB.Write(cfg.Web+"/nodes/", n.IP, n)
				}

				jdb.DB.Write(cfg.Web+"/bitnodes/", bitnode.IP, bitNode)

				fmt.Println("--------------------")
				fmt.Println("bitNodes", bitNodes)
				fmt.Println("--------------------")

				bn.Coin = coin
				bn.BitNodes = append(bn.BitNodes, *bitNode)
			}
			bns = append(bns, bn)
			nodes := jdb.ReadData(cfg.Web + "/nodes")
			ns := make(Nodes, len(nodes))

			for i := range nodes {
				if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
					fmt.Println("Error", err)
				}
			}
			jdb.DB.Write(cfg.Web+"info", "nodes", ns)
		}
	}

	jdb.DB.Write(cfg.Web+"info", "bitnoded", b)
	jdb.DB.Write(cfg.Web+"info", "bitnodestat", bns)

}
