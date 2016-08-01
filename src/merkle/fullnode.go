package merkle

type FullBitcoinNode struct {
	disHonest bool
	block     Block
}

func NewFullBitcoinNode() (node *FullBitcoinNode) {
	node = &FullBitcoinNode{}
	return
}

func (node *FullBitcoinNode) GoDishonest() {
	node.disHonest = true
}
