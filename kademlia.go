package kademlia

import "fmt"

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

type RPCHeader struct {
	Sender    *Contact
	NetworkId string
}

func (k *Kademlia) HandleRPC(request, response *RPCHeader) os.Error {
	if request.NetworkId != k.NetworkId {
		return os.NewError(fmt.Sprintf("Expected network ID %s, got %s", k.NetworkId, request.NetworkId))
	}
	if request.Sender != nil {
		k.routes.Update(request.Sender)
	}
	response.Sender = &k.routes.node
	return nil
}

type KademliaCore struct {
	kad *Kademlia
}

type PingRequest struct {
	RPCHeader
}

type PingResponse struct {
	RPCHeader
}

func (kc *KademliaCore) Ping(args *PingRequest, response *PingResponse) (err os.Error) {
	if err = kc.kad.HandleRPC(&args.RPCHeader, &response.RPCHeader); err == nil {
		log.Stderr("Ping from %s\n", args.RPCHeader)
	}
	return
}
