package merkle

import ()

type MerkleTree struct {
	rows Rows
}

type Rows map[int]*Row // indexed on row index, with leaf row being zero

type Row []*TreeNode

type TreeNode struct {
	hash   []byte
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

func NewTreeNode(left *TreeNode, right *TreeNode) (node *TreeNode) {
	node = &TreeNode{
		hash:  ConcatenateAndHash(left.hash, right.hash),
		left:  left,
		right: right}
	return
}

reconcile having new tree node as well as literal construction

func (tree *MerkleTree) Build(records []*Record) {
	tree.rows = Rows{}
	tree.installStarterRow(records)
	tree.growRowsUpwards()
}

func (tree *MerkleTree) installStarterRow(records []*Record) {
	row := Row{}
	for _, record := range records {
		treeNode := &TreeNode{hash: HashOf(record)}
		row = append(row, treeNode)
	}
	tree.rows[0] = &row
}

func (tree *MerkleTree) growRowsUpwards() {
	for inputRowIdx := 0; !tree.reachedRoot(inputRowIdx); inputRowIdx++ {
		tree.makeRowEvenLengthIfNecessary(inputRowIdx)
		lower := *tree.rows[inputRowIdx]
		upper := Row{}
		for i := 0; i <= len(lower)-2; i += 2 {
			left := lower[i]
			right := lower[i+1]
			newNode := NewTreeNode(left, right)
			left.parent = newNode
			right.parent = newNode
			upper = append(upper, newNode)
		}
		tree.rows[inputRowIdx+1] = &upper
	}
}

func (tree *MerkleTree) reachedRoot(level int) bool {
	return len(*tree.rows[level]) == 1
}

func (tree *MerkleTree) makeRowEvenLengthIfNecessary(level int) {
	row := *tree.rows[level]
	if len(row)%2 == 0 {
		return
	}
	hashToCopy := row[len(row)-1].hash
	row = append(row, &TreeNode{hash: hashToCopy})
	tree.rows[level] = &row
}
