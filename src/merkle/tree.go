package merkle

import ()

type MerkleTree struct {
	rows Rows
}

type Rows map[int] *Row // indexed on row index, with leaf row being zero

type Row []*TreeNode

type TreeNode struct {
    hash    int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

func (tree *MerkleTree) Build(records []Record) {
    tree.rows = Rows{}
	tree.installStarterRow(records)
	tree.growRowAboveLevel(0) // recursive
}

func (tree *MerkleTree) installStarterRow(records []Record) {
	row := Row{}
	for _, record := range records {
		treeNode := &TreeNode{hash: HashOf(record)}
        row = append(row, treeNode)
	}
	tree.rows[0] = &row
}

func (tree *MerkleTree) growRowAboveLevel(levelBelow int) {
    rowBelow := *tree.rows[levelBelow]
    // Reached top?
    if len(rowBelow) == 1 {
        return
    }

    // Augment the row below to make it have even length if neccessary
    if len(rowBelow) % 2 != 0 { 
        hashToCopy := rowBelow[len(rowBelow) - 1].hash
        rowBelow = append(rowBelow, &TreeNode{hash: hashToCopy})
        tree.rows[levelBelow] = &rowBelow
    }

    // Install the row above
    rowAbove := Row{}
    for i := 0; i <= len(below) -2 ; i += 2 {
        left := below[i]
        right := below[i + 1]
        nodeAbove := &TreeNode{
            hash: ConcatenateAndHash(left.hash, right.hash),
            left: left,
            right: right,
        }
        left.parent = nodeAbove
        right.parent = nodeAbove
        rowAbove = append(rowAbove, nodeAbove)
    }
    levelAbove = levelBelow + 1
    tree.Rows[levelAbove] = rowAbove

    // And continue recursively
    tree.installRowAboveLevel(levelAbove)
}

/*
hash for each leaf is int

int hashes themselves stored in rows
each row is a slice
rows stacked in slice

nav....

map keyed on hash
rhs are nav records

nav record gives: left and right children and parent

entry point from pov of leaf is map of record ptr to nav record

type names...

    Row
    TreeNode {left, right, parent}

storage

    recordToTreeNode
    rows
*/
