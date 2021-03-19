package n

import (
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jorm/a"

	"github.com/oknors/okno/pkg/utl"
)

// GetBitNodeStatus returns the full set of information about a node
func GetBitNodeStatus(b a.BitNode) (bitnodeStatus *BitNodeStatus) {
		b.Jrc = utl.NewClient(cfg.CONFIG.RPC.Username, cfg.CONFIG.RPC.Password, b.IP, b.Port)
		if b.Jrc != nil {
			//fmt.Println("bitb:", b.IP)
			getInfo := b.GetInfo()
			getPeerInfo := b.GetPeerInfo()
			getRawMemPool := b.GetRawMemPool()
			getMiningInfo := b.GetMiningInfo()
			getNetworkInfo := b.GetNetworkInfo()

			if getInfo == nil && getPeerInfo == nil && getRawMemPool == nil && getMiningInfo == nil && getNetworkInfo == nil {
				bitnodeStatus = &BitNodeStatus{
					Live: false,
				}
			} else {
				bitnodeStatus = &BitNodeStatus{
					Live:           true,
					IP:             b.IP,
					GetInfo:        getInfo,
					GetPeerInfo:    getPeerInfo,
					GetRawMemPool:  getRawMemPool,
					GetMiningInfo:  getMiningInfo,
					GetNetworkInfo: getNetworkInfo,
					GeoIP:          GetGeoIP(b.IP),
				}
			}
		}
	//fmt.Println("bitnode", bitnode)

	return
}

// GetNodes returns the peers connected to a
func GetNodes(n *BitNodeStatus) (nodes []NodeInfo) {
	// fmt.Println("peers4", peers)
	switch p := n.GetPeerInfo.(type) {
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
