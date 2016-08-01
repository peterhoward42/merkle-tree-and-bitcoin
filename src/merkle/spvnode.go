package merkle

import ()

type SpvBitcoinNode struct {
	remote *FullBitcoinNode
}

func NewSpvBitcoinNode(remote *FullBitcoinNode) *SpvBitcoinNode {
	spvNode := &SpvBitcoinNode{remote: remote}
	return spvNode
}
