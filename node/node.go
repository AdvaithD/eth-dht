package node

import (
    "encoding/hex"
)

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


