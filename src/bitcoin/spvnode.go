package bitcoin

import ()

type SpvBitcoinNode struct {
	remote *FullBitcoinNode
}

func NewSpvBitcoinNode(fullNode *FullBitcoinNode) (node *SpvBitcoinNode) {
	node = &SpvBitcoinNode{remote: fullNode}
	return
}

func (spvNode SpvBitcoinNode) GetAndValidateRecord42FromRemoteNode() {
	record, merklePath := spvNode.remote.GetRecord42()
}
