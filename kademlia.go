package kademlia

type Kademlia struct {
	routes    *RoutingTable
	NetworkId string
}

func NewKademlia(self *Contact, networkId string) (ret *Kademlia) {
	ret = new(Kademlia)
	ret.routes = NewRoutingTable(self)
	ret.NetworkId = network
	return
}
