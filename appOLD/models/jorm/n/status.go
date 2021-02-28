package n

import (
	"github.com/oknors/okno/appOLD/models/jorm/a"

	"github.com/oknors/okno/pkg/utl"
)

// GetBitNodeStatus returns the full set of information about a node
func GetBitNodeStatus(coin string) (bitnode *BitNodeStatus) {
	var live bool
	getInfo := a.RPCSRC(coin).GetInfo()
	getPeerInfo := a.RPCSRC(coin).GetPeerInfo()
	getRawMemPool := a.RPCSRC(coin).GetRawMemPool()
	getMiningInfo := a.RPCSRC(coin).GetMiningInfo()
	getNetworkInfo := a.RPCSRC(coin).GetNetworkInfo()

	if getInfo == nil && getPeerInfo == nil && getRawMemPool == nil && getMiningInfo == nil && getNetworkInfo == nil {
		live = false
	} else {
		live = true
	}

	bitnode = &BitNodeStatus{
		Live:           live,
		GetInfo:        a.RPCSRC(coin).GetInfo(),
		GetPeerInfo:    a.RPCSRC(coin).GetPeerInfo(),
		GetRawMemPool:  a.RPCSRC(coin).GetRawMemPool(),
		GetMiningInfo:  a.RPCSRC(coin).GetMiningInfo(),
		GetNetworkInfo: a.RPCSRC(coin).GetNetworkInfo(),
		GeoIP:          GetGeoIP(a.RPCSRC(coin).IP),
	}

	//fmt.Println("bitnode", bitnode)

	return
}

// GetNodes returns the peers connected to a
func GetNodes(n a.BitNode) (nodes []NodeInfo) {
	peers := n.GetPeerInfo()
	// fmt.Println("peers4", peers)
	switch p := peers.(type) {
	case []interface{}:
		for _, nodeRaw := range p {
			nod := nodeRaw.(map[string]interface{})
			ip, _ := utl.GetIP(nod["addr"].(string))
			// p, _ := strconv.ParseInt(port, 10, 64)
			// n.IP = ip
			// n.Port = p
			node := GetGeoIP(ip)
			// if node != nil {

			nodes = append(nodes, node)
			// }
			// fmt.Println("peersINTERface", nodes)
		}
	}
	return
}
