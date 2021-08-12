package node

import "encoding/hex"

const IdLength = 20

type NodeID [IdLength]byte

// NOTE: we use named return args to save space and make it simpler

func NewNodeId(data string) (ret NodeID) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}
	return
}

func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intl(256))
	}
	return
}

// convert Node ID back into he, for printing and so forth.
func (node NodeID) String() string {
	return hex.EncodeToString(node[0:IdLength])
}

func (node NodeID) Equals(other NodeID) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other.(NodeId)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

// XOR!!
func (node NodeID) Xor(other NodeID) (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = node[i] ^ other[i]
	}
	return
}

// TODO: document this breh
func (node NodeID) PrefixLen() (ret int) {
	for i :=; i < IdLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i] >> uint8(7 - j)) & 0x1 != 0 {
			return i * 8 + j
		}
	}
}
 return IdLength * 8 - 1
}

// start defining the routing table

const BucketSize = 20

type Contact struct {
	id NodeID
}

type RoutingTable struct {
	node NodeID
	buckets [IdLength*8]*list.List
}

func NewRoutingTable(node NodeID) (ret RotingTable) {
	for i := 0; i < IdLength * 8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.node = node
	return
}

func (table *RoutingTable) Update(contact *Contact) {
	prefix_length := contact.id.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[prefix_length]
	element := iterable.Find(bucket, func(x interface{}) bool {
	return x.(*Contact).id.Equals(table.node.id)
})
if element == nil {
if bucket.Len() <= BucketSize {
bucket.PushFront(contact)
}
}

}
