package main

import (
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/bitcoin"
)

/* This program demonstrates the use of Merkle Trees by Bitcoin
 * Single-Payment-Verification (SPV) nodes as part of their strategy to detect
 * dishonest or corrupted replies from their peer-to-peer connections.
 *
 * For a comprehensive written description, primer and background, see the
 * sister article in the docs directory of the source code repository.
 *
 * The code aims to explain and illustrate the operation of Merkle Trees for
 * educational purposes. It exemplifies thoroughly the principles of Merkle
 * Tree operations as used by the Bitcoin protocol, but without the very
 * great complexity of being inside the real system.
 *
 * So we have an object to represent a SPV local node, and an object to
 * represent a full remote node - which holds a full copy of the Blockchain.
 * We illustrate the real protocol's message passing using method calls from
 * the SPV object to the full node object.
 *
 * The full node is coded to hold a mock Blockchain which it builds by sucking
 * the text of some Sherlock Holmes novels off the internet to make each block.
 * Where the real Blockchain's blocks use transactions as the internal records,
 * our blocks split the books' text into 512 byte chunks to simulate
 * individual transactions.
 *
 * The SPV node chooses a block number and the record inside that block it
 * wants to fetch, and initiates the fetching process by requestiong the
 * corresponding block header from its remote full node. When this is
 * returned, the SPV node captures the required Merkle Root hash value which is
 * available in the Block Header.
 *
 * Then it requests the record of interest, and expects the remote node to
 * supply not only the data, but also the Merkle Path that can be used to prove
 * its legitimacy.
 *
 * It then traverses this supplied Merkle Path, so as to calculate the Merkle
 * Root independently, and when this matches the one fetched from the block
 * header, it accepts the record, and treats us by outputting the fragment
 * of Sherlock's text from the record. Otherwise it emits a validation
 * failure message.
 *
 * You are reading the main.go module for the program - which is very short as
 * once it has constructed objects to represent the local and remote node, it
 * delegates all the work to them. You can find the rest of the code in the sub
 * packages: bitcoin, merkle and hash.
 *
 * If you'd like to run the code yourself, you can find instructions about how
 * to install Go here: https://golang.org/doc/install
 *
 * Downloading and using code from Github repositories like this is an
 * intrinsic part of the Go language - so the instructions above will also show
 * you how to get and run this demo code.
 */

func main() {

	// Real SPV nodes have to discover some remote full nodes to talk to.
	// We make life a little simpler with an arranged marriage...

	remoteNode := bitcoin.NewFullBitcoinNode()
	localNode := bitcoin.SpvBitcoinNode{Remote: &remoteNode}

	blockOfInterest := 2
	recordToFetch := 42

	record, err := localNode.FetchAndValidateRecordFromRemote(
		blockOfInterest, recordToFetch)

	if err != nil {
		fmt.Printf("Fetch failed with: %v", err)
	} else {
		fmt.Printf("Fetched ths record which passed validation:\n\n%s\n",
			record)
	}
}
