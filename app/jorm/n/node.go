package n

import (
	"encoding/json"
	"fmt"
	"github.com/oknors/okno/app/jorm/a"

	"github.com/oknors/okno/app/jorm/c"
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/pkg/utl"
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
	IP	           string        `json:"ip"`
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
		www := cfg.Path + "/jorm/conf/" + coin.Slug + "/"
		if utl.FileExists(www + "bitnodes") {
			b = append(b, coin.Slug)
			bitNodes := a.BitNodes{}
			if err := jdb.JDB.Read("jorm/conf/"+coin.Slug, "bitnodes", &bitNodes); err != nil {
				fmt.Println("Error", err)
			}
			for _, bitnode := range bitNodes {
				bitNode := GetBitNodeStatus(bitnode)
				nds := GetNodes(bitNode)
				for _, n := range nds {
				jdb.JDB.Write("jorm/nodes/", n.IP, n)
				}

				jdb.JDB.Write("jorm/bitnodes/"+coin.Slug, bitnode.IP, bitNode)
				//
				//fmt.Println("--------------------")
				//fmt.Println("bitNodes", nds)
				//fmt.Println("--------------------")

				bn.Coin = coin
				bn.BitNodes = append(bn.BitNodes, *bitNode)
			}
			bns = append(bns, bn)

			data, err := jdb.JDB.ReadAll("jorm/nodes")
				utl.ErrorLog(err)
			nodes := make([][]byte, len(data))
				for i := range data {
					nodes[i] = []byte(data[i])
				}


			ns := make(Nodes, len(nodes))
			//
			for i := range nodes {
				if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
					fmt.Println("Error", err)
				}
			}
			jdb.JDB.Write("jorm/info", "nodes", ns)
		}
	}

	jdb.JDB.Write("jorm/info", "bitnoded", b)
	jdb.JDB.Write("jorm/info", "bitnodestat", bns)

}
