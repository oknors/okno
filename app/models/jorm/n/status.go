package n

import (
	"github.com/oknors/okno/app/models/jorm/a"

	"github.com/oknors/okno/app/utl"
)

// GetBitNodeStatus returns the full set of information about a node
func GetBitNodeStatus(a a.BitNode) (bitnode *BitNodeStatus) {
	var live bool
	getInfo := a.GetInfo()
	getPeerInfo := a.GetPeerInfo()
	getRawMemPool := a.GetRawMemPool()
	getMiningInfo := a.GetMiningInfo()
	getNetworkInfo := a.GetNetworkInfo()

	if getInfo == nil && getPeerInfo == nil && getRawMemPool == nil && getMiningInfo == nil && getNetworkInfo == nil {
		live = false
	} else {
		live = true
	}

	bitnode = &BitNodeStatus{
		Live:           live,
		GetInfo:        a.GetInfo(),
		GetPeerInfo:    a.GetPeerInfo(),
		GetRawMemPool:  a.GetRawMemPool(),
		GetMiningInfo:  a.GetMiningInfo(),
		GetNetworkInfo: a.GetNetworkInfo(),
		GeoIP:          GetGeoIP(a.IP),
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
