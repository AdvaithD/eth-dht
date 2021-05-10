

import "encoding/hex"

const IdLength = 20

type NodeID [IdLength]byte

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
